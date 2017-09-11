package oaicrawl

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/sethgrid/pester"
	log "github.com/sirupsen/logrus"
)

// Harvester encapsulates harvesting options.
type Harvester struct {
	Base           string
	Format         string
	MaxRetries     int
	MaxElapsedTime time.Duration
	NumWorkers     int
	Verbose        bool
	BestEffort     bool
	Output         io.Writer

	wg      sync.WaitGroup
	queue   chan work
	results chan result
	done    chan bool
}

// NewHarvester creates a new harvester for an endpoint with default options.
func NewHarvester(base string) *Harvester {
	return &Harvester{
		Base:           base,
		Format:         "oai_dc",
		MaxElapsedTime: 10 * time.Second,
		MaxRetries:     3,
		NumWorkers:     4 * runtime.NumCPU(),
		Output:         os.Stdout,
	}
}

type work struct {
	Identifier string
}

type result struct {
	Body []byte
	Err  error
}

// worker takes an item of the queue of work items, fetches the content, retries
// on various errors and sends the result to the output.
func (h *Harvester) worker(name string) {
	defer h.wg.Done()

	log.Debug(name, " started")

	client := pester.New()
	client.Timeout = 5 * time.Second
	client.MaxRetries = h.MaxRetries
	client.Backoff = pester.ExponentialBackoff
	client.LogHook = func(e pester.ErrEntry) {
		s := e.URL
		if len(s) > 45 {
			s = ".." + s[len(s)-45:]
		}
		log.Warn(name, " backoff [", e.Attempt, "]: ", s)
	}

	var i int
	for item := range h.queue {
		link := fmt.Sprintf("%s?verb=GetRecord&identifier=%s&metadataPrefix=%s",
			h.Base, item.Identifier, h.Format)

		op := func() error {
			// Fetch link.
			resp, err := client.Get(link)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			// Check for OAI protocol errors.
			var generic GenericResponse
			dec := xml.NewDecoder(bytes.NewReader(b))
			dec.Strict = false
			if err := dec.Decode(&generic); err != nil {
				return err
			}
			if generic.Error.Code != "" {
				switch generic.Error.Code {
				// Do not treat missing id as an error.
				case "idDoesNotExist":
					log.Debug("skipping id ", item.Identifier)
					return nil
				default:
					return fmt.Errorf("%s oai error [%s]: %s %s", link,
						name, generic.Error.Code, generic.Error.Message)
				}
			}

			h.results <- result{Body: b}

			i++
			if i%100 == 0 {
				log.Debug(name, " completed ", i, " requests")
			}
			return nil
		}

		// Retry op on HTTP, XML decoding or oai protocol errors.
		eb := backoff.NewExponentialBackOff()
		eb.MaxElapsedTime = h.MaxElapsedTime
		err := backoff.RetryNotify(op, eb, func(err error, _ time.Duration) {
			log.Warn(fmt.Sprintf("%s retry reason: %s", name, err))
		})

		// Finally, if we still encounter an error, report it.
		if err != nil {
			h.results <- result{Err: err}
		}
	}
	log.Debug(name, " shut down")
}

// write collects data from the output channel and writes it to the configured writer.
func (h *Harvester) write() {
	var i int
	for r := range h.results {
		if r.Err != nil {
			if h.BestEffort {
				log.Warn(r.Err)
			} else {
				log.Fatal(r.Err)
			}
		}
		if _, err := h.Output.Write(r.Body); err != nil {
			log.Fatal(err)
		}
		i++
		if i%1000 == 0 {
			log.Debug("writer: written ", i, " records")
		}
	}
	h.done <- true
}

// Run starts the harvest with the given parameters.
func (h *Harvester) Run() error {
	started := time.Now()

	h.queue = make(chan work)
	h.results = make(chan result)
	h.done = make(chan bool)

	for i := 0; i < h.NumWorkers; i++ {
		h.wg.Add(1)
		go h.worker(fmt.Sprintf("worker-%02d", i))
	}

	go h.write()

	link := fmt.Sprintf("%s?verb=ListIdentifiers&metadataPrefix=%s", h.Base, h.Format)
	var items, requests int

	client := pester.New()
	client.MaxRetries = h.MaxRetries
	client.Backoff = pester.ExponentialBackoff
	client.LogHook = func(e pester.ErrEntry) {
		log.Warn("main client: ", e)
	}

	for {
		log.Debug(link)
		resp, err := client.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		requests++
		var lir ListIdentifiersResponse

		dec := xml.NewDecoder(resp.Body)
		dec.Strict = false
		if err := dec.Decode(&lir); err != nil {
			log.Fatal(err)
		}
		for _, item := range lir.ListIdentifiers.Headers {
			h.queue <- work{Identifier: item.Identifier}
			items++
		}
		token := lir.ListIdentifiers.ResumptionToken
		if token.Value == "" {
			break
		}
		link = fmt.Sprintf("%s?verb=ListIdentifiers&resumptionToken=%s", h.Base, token.Value)
		if requests%10 == 0 {
			log.Debug("completed ", requests, " ListIdentifier requests ",
				items, "/", token.Cursor, "/", token.CompleteListSize)
		}
	}

	log.Debug("shutting down workers")

	close(h.queue)
	h.wg.Wait()
	close(h.results)
	<-h.done

	log.Debug("fetched ", items, " identifiers with ",
		requests, " requests in ", time.Since(started))

	return nil
}

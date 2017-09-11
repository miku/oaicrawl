// Example for looping over ListIdentifiers responses.
//
//     $ go run examples/listids/main.go
//     oai:www.duo.uio.no:10852/28742
//     oai:www.duo.uio.no:10852/28743
//     ...
//     oai:www.duo.uio.no:10852/9316
//     oai:www.duo.uio.no:10852/9622
//     Note: There are more identifiers (45828)
//
package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/miku/haprot"
)

func main() {
	started := time.Now()
	// base := "http://www.duo.uio.no/oai/request"
	// base := "http://zvdd.de/oai2"
	// base := "http://oai.narcis.nl/oai"
	base := "http://www.digizeitschriften.de/oai2/"
	link := fmt.Sprintf("%s?verb=ListIdentifiers&metadataPrefix=oai_dc", base)
	var items, requests int
	var token string
	for {
		resp, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		requests++
		var lir haprot.ListIdentifiersResponse
		if err := xml.NewDecoder(resp.Body).Decode(&lir); err != nil {
			log.Fatal(err)
		}
		for _, item := range lir.ListIdentifiers.Headers {
			fmt.Println(item.Identifier)
			items++
		}
		if token = lir.ListIdentifiers.ResumptionToken.Value; token == "" {
			break
		}
		link = fmt.Sprintf("%s?verb=ListIdentifiers&resumptionToken=%s", base, token)
	}
	fmt.Fprintf(os.Stderr, "fetched %d identifiers with %d requests in %s\n",
		items, requests, time.Since(started))
}

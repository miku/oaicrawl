// oaicrawl downloads a complete endpoint by requesting records one by one.
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/miku/oaicrawl"
	log "github.com/sirupsen/logrus"
)

// ApplicationVersion is the application version.
const ApplicationVersion = "0.1.0"

var (
	maxRetries     = flag.Int("retry", 3, "max number of retries")
	format         = flag.String("f", "oai_dc", "format")
	verbose        = flag.Bool("verbose", false, "more logging")
	version        = flag.Bool("version", false, "show version")
	bestEffort     = flag.Bool("b", false, "create best effort data set")
	maxElapsedTime = flag.Duration("e", 12*time.Second, "max elapsed time")
	numWorkers     = flag.Int("w", 4*runtime.NumCPU(), "number of parallel connections")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Println(ApplicationVersion)
		os.Exit(0)
	}

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	if flag.NArg() == 0 {
		log.Fatal("endpoint required")
	}

	if *verbose {
		log.SetLevel(log.DebugLevel)
	}

	harvester := oaicrawl.NewHarvester(flag.Arg(0))
	harvester.MaxRetries = *maxRetries
	harvester.MaxElapsedTime = *maxElapsedTime
	harvester.Format = *format
	harvester.BestEffort = *bestEffort
	harvester.NumWorkers = *numWorkers

	if err := harvester.Run(); err != nil {
		log.Fatal(err)
	}
}

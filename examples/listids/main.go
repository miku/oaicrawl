// Example for ListIdentifiers response.
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

	"github.com/miku/oaicrawl"
)

func main() {
	resp, err := http.Get("http://www.duo.uio.no/oai/request?verb=ListIdentifiers&metadataPrefix=oai_dc")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var lir oaicrawl.ListIdentifiersResponse
	if err := xml.NewDecoder(resp.Body).Decode(&lir); err != nil {
		log.Fatal(err)
	}
	for _, item := range lir.ListIdentifiers.Headers {
		fmt.Printf("%s\n", item.Identifier)
	}
	if lir.ListIdentifiers.ResumptionToken.Value == "" {
		fmt.Fprintf(os.Stderr, "Note: These are all available identifiers.")
		os.Exit(0)
	} else {
		fmt.Fprintf(os.Stderr, "Note: There are more identifiers (%s/%s)\n",
			lir.ListIdentifiers.ResumptionToken.Cursor,
			lir.ListIdentifiers.ResumptionToken.CompleteListSize)
	}

	// If there are more, assemble the next URL.
	token := lir.ListIdentifiers.ResumptionToken.Value
	next := fmt.Sprintf("http://www.duo.uio.no/oai/request?verb=ListIdentifiers&resumptionToken=%s", token)

	resp, err = http.Get(next)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if err := xml.NewDecoder(resp.Body).Decode(&lir); err != nil {
		log.Fatal(err)
	}
	for _, item := range lir.ListIdentifiers.Headers {
		fmt.Printf("%s\n", item.Identifier)
	}
	if lir.ListIdentifiers.ResumptionToken.Value == "" {
		fmt.Fprintf(os.Stderr, "Note: These are all available identifiers.")
	} else {
		fmt.Fprintf(os.Stderr, "Note: There are more identifiers (%s/%s)\n",
			lir.ListIdentifiers.ResumptionToken.Cursor,
			lir.ListIdentifiers.ResumptionToken.CompleteListSize)
	}
}

// Example for Idenfify response.
//
//     $ go run examples/identify/main.go
//     response date    2017-09-11T14:20:10Z
//     name             arXiv
//     url              http://export.arxiv.org/oai2
//     version          2.0
//     admin            [help@arxiv.org]
//     earliest date    2007-05-23
//     delete policy    persistent
//     granularity      YYYY-MM-DD
//
package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/miku/haprot"
)

func main() {
	resp, err := http.Get("http://export.arxiv.org/oai2?verb=Identify")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var ir haprot.IdentifyResponse
	if err := xml.NewDecoder(resp.Body).Decode(&ir); err != nil {
		log.Fatal(err)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	defer w.Flush()
	fmt.Fprintf(w, "response date\t%s\n", ir.ResponseDate)
	fmt.Fprintf(w, "name\t%s\n", ir.Identify.RepositoryName)
	fmt.Fprintf(w, "url\t%s\n", ir.Identify.BaseURL)
	fmt.Fprintf(w, "version\t%s\n", ir.Identify.ProtocolVersion)
	fmt.Fprintf(w, "admin\t%s\n", ir.Identify.AdminEmail)
	fmt.Fprintf(w, "earliest date\t%s\n", ir.Identify.EarliestDatestamp)
	fmt.Fprintf(w, "delete policy\t%s\n", ir.Identify.DeletedRecord)
	fmt.Fprintf(w, "granularity\t%s\n", ir.Identify.Granularity)
}

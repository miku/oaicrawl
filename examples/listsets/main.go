// Example for ListSets response.
//
//     $ go run examples/listsets/main.go
//     Computer Science (cs)
//     Mathematics (math)
//     Physics (physics)
//     Astrophysics (physics:astro-ph)
//     Condensed Matter (physics:cond-mat)
//     General Relativity and Quantum Cosmology (physics:gr-qc)
//     High Energy Physics - Experiment (physics:hep-ex)
//     High Energy Physics - Lattice (physics:hep-lat)
//     High Energy Physics - Phenomenology (physics:hep-ph)
//     High Energy Physics - Theory (physics:hep-th)
//     Mathematical Physics (physics:math-ph)
//     Nonlinear Sciences (physics:nlin)
//     Nuclear Experiment (physics:nucl-ex)
//     Nuclear Theory (physics:nucl-th)
//     Physics (Other) (physics:physics)
//     Quantum Physics (physics:quant-ph)
//     Quantitative Biology (q-bio)
//     Quantitative Finance (q-fin)
//     Statistics (stat)
//
package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/miku/haprot"
)

func main() {
	resp, err := http.Get("http://export.arxiv.org/oai2?verb=ListSets")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var lsr haprot.ListSetsResponse
	if err := xml.NewDecoder(resp.Body).Decode(&lsr); err != nil {
		log.Fatal(err)
	}
	var sets []string
	for _, s := range lsr.ListSets.Sets {
		sets = append(sets, fmt.Sprintf("%s (%s)", s.SetName, s.SetSpec))
	}
	fmt.Println(strings.Join(sets, "\n"))
	if lsr.ListSets.ResumptionToken.Value == "" {
		fmt.Println("Note: These are all available sets.")
	} else {
		fmt.Println("Note: There are more sets.")
	}
}

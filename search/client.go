package search

import (
	"log"
	"sync"
)

//Result represents the high level interface which every result from differnt providers adhere to.
type Result interface {
	PackagePath() string
	Info() string
}

//Search retuns a slice of Result and an error if any one of the search providers failed or timed out.
func Search(query string) ([]Result, error) {
	results := make(chan Result)
	ech := make(chan error)
	done := make(chan bool)
	resultsMap := make(map[string]bool)
	var wg sync.WaitGroup
	wg.Add(2)
	var Results []Result
	go searchGoDoc(query, results, ech, &wg)
	go searchGoSearch(query, results, ech, &wg)
	go func() {
		wg.Wait()
		done <- true
		close(results)
		close(ech)
	}()
	for {
		select {
		case values := <-results:
			if values == nil {
				log.Println("Recieved a nil info for", values)
				break
			}
			if _, ok := resultsMap[values.PackagePath()]; !ok {
				resultsMap[values.PackagePath()] = true
				Results = append(Results, values)
			}
		case <-done:
			return Results, nil
		case err := <-ech:
			return nil, err
		}
	}
}

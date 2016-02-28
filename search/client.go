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
func Search(query string, maxCount int) ([]Result, error) {
	results := make(chan Result)
	ech := make(chan error)
	done := make(chan bool)
	terminate := make(chan bool, 1)
	resultsMap := make(map[string]bool)
	var wg sync.WaitGroup
	wg.Add(2)
	var resultItems []Result
	go searchGoDoc(query, results, terminate, ech, &wg)
	go searchGoSearch(query, results, terminate, ech, &wg)
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
				if len(values.Info()) > 3 { //Not considering packages with description that are not realistic
					resultsMap[values.PackagePath()] = true
					resultItems = append(resultItems, values)
					if len(resultItems) > maxCount {
						terminate <- true
						terminate <- true
						return resultItems, nil
					}
				}
			}
		case <-done:
			return resultItems, nil
		case err := <-ech:
			return nil, err
		}
	}
}

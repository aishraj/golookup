package main

import (
	"fmt"
	"log"
	"sync"
)

type searchResult interface {
	PackagePath() string
	Info() string
}

type goDocResult struct {
	Path     string `json:"path"`
	Synopsis string `json:"synopsis"`
}

type goDocResponse struct {
	Results []goDocResult `json:"results"`
}

type goSearchResult struct {
	Name        string `json:"name"`
	Package     string `json:"package"`
	Author      string `json:"author"`
	Synopsis    string `json:"synopsis"`
	Description string `json:"description"`
	ProjectURL  string `json:"projecturl"`
}

func (result goDocResult) PackagePath() string {
	return result.Path
}

func (result goDocResult) Info() string {
	return result.Synopsis
}

func (result goSearchResult) PackagePath() string {
	return result.ProjectURL
}

func (result goSearchResult) Info() string {
	return result.Description
}

func search(searchTerm string) {
	results := make(chan searchResult)
	go findResults(searchTerm, results)
	for result := range results {
		fmt.Printf("%v : %v\n", result.PackagePath(), result.Info())
	}
}

func findResults(term string, ch chan<- searchResult) {
	dummyResult := goDocResult{"github.com/aishraj/gohort", "A silly library"}
	ch <- dummyResult
	close(ch)
}

type resultSet struct {
	Results map[string]bool
	sync.RWMutex
}

func main() {
	results := make(chan searchResult)
	ech := make(chan error)
	done := make(chan bool)
	go searchGoDoc("rss", results, ech, done)

	for {
		select {
		case err := <-ech:
			log.Fatal("Encountered error", err)
			return
		case result := <-results:
			log.Println(result)
		case <-done:
			log.Println("Done")
			return
		}
	}
}

// func main() {
//
// }

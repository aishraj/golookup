package main

import (
	"fmt"
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

type goSearchResponse struct {
	Query string           `json:"query"`
	Hits  []goSearchResult `json:"hits"`
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

func main() {
	results := make(chan searchResult)
	ech := make(chan error)
	resultsMap := make(map[string]bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go searchGoDoc("rss", results, ech, &wg)
	go searchGoSearch("rss", results, ech, &wg)
	go func() {
		wg.Wait()
		close(results)
		close(ech)
	}()
	for values := range results {
		if _, ok := resultsMap[values.PackagePath()]; !ok {
			resultsMap[values.PackagePath()] = true
			fmt.Printf("%v\n%v\n\n", values.PackagePath(), values.Info())
		}
	}
}

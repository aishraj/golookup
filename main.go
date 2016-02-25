package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	return result.Synopsis
}

func (result goSearchResult) String() string {
	return fmt.Sprintf("%v\n%v\n\n", result.PackagePath(), result.Info())
}

func (result goDocResult) String() string {
	return fmt.Sprintf("%v\n%v\n\n", result.PackagePath(), result.Info())
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gofind [search term]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	if len(os.Args) != 2 {
		flag.Usage()
		return
	}
	searchTerm := os.Args[1]
	results, err := search(searchTerm)
	if err != nil {
		log.Panic("Encountered error. ", err)
	}
	if len(results) == 0 {
		fmt.Println("Did not get any results for the given query.")
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

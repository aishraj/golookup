package search

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	goDocSearchURI = "http://api.godoc.org/search?"
)

type goSearchResponse struct {
	Query string           `json:"query"`
	Hits  []goResult `json:"hits"`
}

type goResult struct {
	Name        string `json:"name"`
	Package     string `json:"package"`
	Author      string `json:"author"`
	Synopsis    string `json:"synopsis"`
	Description string `json:"description"`
	ProjectURL  string `json:"projecturl"`
}

func (result goResult) PackagePath() string {
	return result.ProjectURL
}

func (result goResult) Info() string {
	return result.Synopsis
}

func (result goResult) String() string {
	return fmt.Sprintf("%v\n%v\n\n", result.PackagePath(), result.Info())
}

const (
	goSearchURI = "http://go-search.org/api?"
)

func searchGoSearch(term string, ch chan<- Result, ech chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	uriValues := url.Values{}
	uriValues.Set("action", "search")
	uriValues.Set("q", term)
	params := uriValues.Encode()
	urlString := goSearchURI + params
	timeout := time.Duration(1 * time.Minute)
	client := http.Client{Timeout: timeout}
	resp, err := client.Get(urlString)
	if err != nil {
		log.Printf("Encountered an error while trying to do an HTTP GET %v \n", err)
		ech <- err
		return
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var responseBody goSearchResponse
	err = decoder.Decode(&responseBody)
	if err != nil {
		log.Println("Encountered an error while trying to parse JSON response body. Error is ", err)
		ech <- err
		return
	}
	for _, item := range responseBody.Hits {
		ch <- item
	}
}

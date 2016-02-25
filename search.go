package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
)

func searchGoDoc(term string, ch chan<- searchResult, ech chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	uriValues := url.Values{}
	uriValues.Set("q", term)
	parameters := uriValues.Encode()
	urlString := "http://api.godoc.org/search?" + parameters
	resp, err := http.Get(urlString)
	if err != nil {
		log.Println("Encountered an error wihle trying to do an HTTP GET", err)
		ech <- err
		return
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var responseData goDocResponse
	err = decoder.Decode(&responseData)
	if err != nil {
		log.Println("Encountered an error while trying to parse JSON response body. Error is", err)
		ech <- err
		return
	}
	for _, item := range responseData.Results {
		ch <- item
	}
}

func searchGoSearch(term string, ch chan<- searchResult, ech chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	uriValues := url.Values{}
	uriValues.Set("action", "search")
	uriValues.Set("q", term)
	params := uriValues.Encode()
	urlString := "http://go-search.org/api?" + params
	resp, err := http.Get(urlString)
	if err != nil {
		log.Printf("Encountered an error wihle trying to do an HTTP GET %v \n", err)
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

func search(query string) ([]searchResult, error) {
	results := make(chan searchResult)
	ech := make(chan error)
	done := make(chan bool)
	resultsMap := make(map[string]bool)
	var wg sync.WaitGroup
	wg.Add(2)
	var searchResults []searchResult
	go searchGoDoc(query, results, ech, &wg)
	go searchGoSearch(query, results, ech, &wg)
	go func() {
		wg.Wait()
		close(results)
		close(ech)
		done <- true
	}()
	for {
		select {
		case values := <-results:
			if _, ok := resultsMap[values.PackagePath()]; !ok {
				resultsMap[values.PackagePath()] = true
				searchResults = append(searchResults, values)
			}
		case <-done:
			return searchResults, nil
		case err := <-ech:
			return nil, err
		}
	}
}

func main() {
	results, err := search("rss")
	if err != nil {
		log.Panic("Encountered error. ", err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

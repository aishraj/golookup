package main

import (
	"encoding/json"
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
		log.Println("error is ********", err)
		ech <- err
		return
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var responseData goDocResponse
	err = decoder.Decode(&responseData)
	if err != nil {
		log.Println("error is ********", err)
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
		log.Printf("Error is %v \n", err)
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

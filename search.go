package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func searchGoDoc(term string, ch chan<- searchResult, ech chan<- error, done chan<- bool) {
	defer close(ch)
	defer close(ech)
	uriValues := url.Values{}
	uriValues.Set("q", term)
	parameters := uriValues.Encode()
	urlString := "http://api.godoc.org/search?" + parameters
	log.Println("URL is", urlString)
	resp, err := http.Get(urlString)
	if err != nil {
		log.Println("error is ********", err)
		ech <- err
		return
	}
	defer resp.Body.Close()
	//log.Println("response body is ", resp.Body)
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
	done <- true
	return
}

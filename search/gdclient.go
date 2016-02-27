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

type goDocResult struct {
	Path     string `json:"path"`
	Synopsis string `json:"synopsis"`
}

type goDocResponse struct {
	Results []goDocResult `json:"results"`
}

func (result goDocResult) PackagePath() string {
	return result.Path
}

func (result goDocResult) Info() string {
	return result.Synopsis
}

func (result goDocResult) String() string {
	return fmt.Sprintf("%v\n%v\n\n", result.PackagePath(), result.Info())
}

func searchGoDoc(term string, ch chan<- Result, ech chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	uriValues := url.Values{}
	uriValues.Set("q", term)
	parameters := uriValues.Encode()
	urlString := goDocSearchURI + parameters
	timeout := time.Duration(1 * time.Minute)
	client := http.Client{Timeout: timeout}
	resp, err := client.Get(urlString)
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

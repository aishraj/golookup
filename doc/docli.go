package doc

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	docBaseURL = "https://godoc.org/"
)

//FetchDoc fetches the documentation of a given package string from godoc.org.
func FetchDoc(pkgname string) (string, error) {
	timeout := time.Duration(1 * time.Minute)
	client := http.Client{Timeout: timeout}
	req, err := http.NewRequest("GET", docBaseURL+pkgname, nil)
	req.Header.Add("Accept", "text/plain")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Encountered an error while trying to perform GET on package doc. Error is :", err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("Did not get the expected response from the server")
		//TODO: Should we return here ?
	}
	rawData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Unable to read the response body")
		return "", err
	}
	return string(rawData), nil
}

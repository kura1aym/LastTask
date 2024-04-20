package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	apiKey        = "ae073b069fmsh5c57783e333a5dap16ca47jsn559ad372b001"
	apiHost       = "shazam.p.rapidapi.com"
	delayDuration = time.Second * 1
)

var apiURLs = map[string]string{
	"events":       "https://shazam.p.rapidapi.com/shazam-events/list?artistId=73406786&l=en-US&from=2022-12-31&limit=50&offset=0",
	"search":       "https://shazam.p.rapidapi.com/search?term=kiss%20the%20rain&locale=en-US&offset=0&limit=5",
	"autocomplete": "https://shazam.p.rapidapi.com/auto-complete?term=kiss%20the&locale=en-US",
}

type APIResponse map[string]interface{}

func fetchAPI(url string, headers map[string]string, results chan<- APIResponse, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(delayDuration)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Request creation error: %s", err)
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Request execution error: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading the response body: %s", err)
		return
	}

	var responseData APIResponse
	if err := json.Unmarshal(body, &responseData); err != nil {
		log.Printf("JSON parsing error: %s", err)
		return
	}

	results <- responseData
}

func main() {
	results := make(chan APIResponse)
	var wg sync.WaitGroup

	headers := map[string]string{
		"X-RapidAPI-Key":  apiKey,
		"X-RapidAPI-Host": apiHost,
	}

	for _, url := range apiURLs {
		wg.Add(1)
		go fetchAPI(url, headers, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for responseData := range results {
		fmt.Printf("Result: %+v\n", responseData)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Send hash to each worker's overlay ip address.
func sendToWorkers(hash string, workerAddrs []string) *http.Response {
	var workerCount int64 = int64(len(workerAddrs))  // How many slave ips have we registered?
	var pages int64 = dictionaryLength / workerCount // How big will the assigned tasks will be
	myResponse := make(chan *http.Response, 0)

	// Iterate over worker addresses and submit a hash, startindex, and length of entries.
	for i := 0; i < len(workerAddrs); i++ {
		startIndex := int64(i) * pages

		// This go routine submits values to the PODS not the ec2s.
		go func(index int) {

			resp, er := tryGet(workerAddrs[index], hash, startIndex, pages)
			if er == nil {
				myResponse <- resp
				return
			}
			log.Fatalf("Timeout: failed to connect - %v", er)

		}(i)
	}
	return <-myResponse
}

// tryGet attempts to send the hash string to the address every second, for t seconds. If no connection is made in that time, returns an error.
func tryGet(addr, hash string, index int64, length int64) (*http.Response, error) {
	var er error
	var colliderPort string = "8080"

	// Submit request to colliders. Post request will be used because GET query and
	// POST Content-Type: application/x-www-form-urlencoded get parsed exactly the same by
	// myURL.PasreForm()
	colliderURL := fmt.Sprintf("http://%s:%s/", addr, colliderPort)
	stringIndex := string(index)
	stringLength := string(length)
	// contentType := "application/x-www-form-urlencoded"
	// content := fmt.Sprintf("hash=%s&index=%s&length=%s", hash, index, length)

	resp, er := http.PostForm(colliderURL, url.Values{"hash": {hash}, "index": {stringIndex}, "length": {stringLength}})

	// Normal operation
	if er == nil {
		return resp, er
	}

	// We did not connect to collider in time. Don't return a resp back to sendToWorkers.
	// Return error from POST request.
	log.Printf("My error says =====> %v", er)
	return nil, er
}

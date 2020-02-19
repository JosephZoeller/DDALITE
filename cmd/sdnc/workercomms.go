package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Send hash to each worker's ip address.
func sendToWorkers(hash string, workerAddrs []string) *http.Response {
	workerCount := len(workerAddrs)                // How many slave ips have we registered?
	var pages int = dictionaryLength / workerCount // How big will the assigned tasks will be

	// Iterate over worker addresses and submit a hash, startindex, and length of entries.
	for i := 0; i < len(workerAddrs); i++ {
		startIndex := i * pages
		go func(index int) {
			resp, er := tryGet(workerAddrs[index], hash, startIndex, pages, 15)
			if er == nil {
				return resp
			}
			log.Fatal("Timeout: failed to connect - ", er)

		}(i)
	}
}

// tryGet attempts to send the hash string to the address every second, for t seconds. If no connection is made in that time, returns an error.
func tryGet(addr, hash string, index int, length int, t int) (*http.Response, error) {
	var er error
	var resp *http.Response
	var colliderPort string = "8080"

	for i := 0; i < t; i++ {
		// Submit request to colliders. Post request will not be used because GET query and
		// POST Content-Type: application/x-www-form-urlencoded get parsed exactly the same by
		// myURL.PasreForm()
		colliderURL := fmt.Sprintf("http://%s:%s/", addr, colliderPort)
		contentType := "application/x-www-form-urlencoded"
		content := fmt.Sprintf("hash=%s&index=%s&length=%s", hash, index, length)
		resp, er = http.Post(colliderURL, contentType, content)
		defer resp.Close()
		// Normal operation
		if er == nil {
			return resp, er
		}
		// Print error to standard console and wait to try again.
		fmt.Println(er)
		time.Sleep(time.Second)
	}
	// We did not connect to collider in time. Don't return a resp back to sendToWorkers.
	// Return error from GET request.
	return nil, er
}

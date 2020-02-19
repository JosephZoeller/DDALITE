package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Send hash to each worker's ip address.
func sendToWorkers(hash string, workerAddrs []string) {
	for i := 0; i < len(workerAddrs); i++ {
		go func(index int) {
			resp, er := tryGet(workerAddrs[index], hash, 10)
			if er == nil {
				resp.Write(os.Stdout)
			} else {
				log.Fatal("Timeout: failed to connect - ", er)
			}
		}(i)
	}
}

// tryGet attempts to send the hash string to the address every second, for t seconds. If no connection is made in that time, returns an error.
func tryGet(addr, hash string, t int) (*http.Response, error) {
	var er error
	var resp *http.Response
	var colliderPort string = "8080"

	for i := 0; i < t; i++ {
		// Submit Get request to colliders. Post request will not be used because GET query and
		// POST Content-Type: application/x-www-form-urlencoded get parsed exactly the same by
		// myURL.PasreForm()
		resp, er = http.Get(fmt.Sprintf("%s:%s/?hash=%s&start=%s&end=%s", addr, colliderPort, hash))
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

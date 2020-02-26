package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

var collision []cityhashutil.HashCollision

// Send hash to each worker's overlay ip address.
func sendToWorkers(hash string, workerAddrs []string) string {
	responseChan := make(chan io.ReadCloser, len(workerAddrs))
	var workerCount int64 = int64(len(workerAddrs))  // How many slave ips have we registered?
	var pages int64 = dictionaryLength / workerCount // How big will the assigned tasks will be

	var tmp cityhashutil.HashCollision

	fmt.Println("Beginning Sendout Logic. # Worker Addresses: ", len(workerAddrs))
	for i := 0; i < len(workerAddrs); i++ {
		startIndex := int64(i) * pages

		// This go routine submits values to the PODS not the ec2s.
		go func(index int) {

			resp, er := tryGet(workerAddrs[index], hash, startIndex, pages)
			if er != nil {
				log.Println(er)
			}
			fmt.Println("Response from worker: ", resp.Body)
			time.Sleep(time.Second * 5)
			responseChan <- resp.Body

		}(i)
	}

	fmt.Println("Beginning Return Logic")
	for i := 0; i< len(workerAddrs); i++ {
		firstResponseBody := <- responseChan
		fmt.Println("firstResponseBody: ", firstResponseBody)
		json.NewDecoder(firstResponseBody).Decode(&tmp)

		fmt.Println("temp collision: ", tmp.Collision, " | Input hash: ", tmp.InputHash, " | error message: ", tmp.Err )
		if tmp.Collision != "" {
			return tmp.Collision
		}
	}
	return tmp.Collision
}

// tryGet attempts to send the hash string to the address every second, for t seconds. If no connection is made in that time, returns an error.
func tryGet(addr, hash string, index int64, length int64) (*http.Response, error) {
	var er error
	var colliderPort string = "8080"

	// Submit request to colliders. Post request will be used because GET query and
	// POST Content-Type: application/x-www-form-urlencoded get parsed exactly the same by
	// myURL.PasreForm()
	colliderURL := fmt.Sprintf("http://%s:%s/", addr, colliderPort)
	// contentType := "application/x-www-form-urlencoded"
	// content := fmt.Sprintf("hash=%s&index=%s&length=%s", hash, index, length)

	request, er := http.NewRequest("GET", colliderURL+"?hash="+hash+"&index="+fmt.Sprintf("%d", index)+"&length="+fmt.Sprintf("%d", length), nil)
	if er != nil {
		log.Println(er, "Backend server connection failed")
	}
	client := http.Client{}
	response, er := client.Do(request)

	return response, er
}

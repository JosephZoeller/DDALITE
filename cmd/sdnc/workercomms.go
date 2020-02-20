package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Data is result struct
type Data struct {
	Hash   string `json:"hash"`
	Result string `json:"result"`
}

var result []Data

// Send hash to each worker's overlay ip address.
func sendToWorkers(hash string, workerAddrs []string) []Data {
	var workerCount int64 = int64(len(workerAddrs))  // How many slave ips have we registered?
	var pages int64 = dictionaryLength / workerCount // How big will the assigned tasks will be

	// Iterate over worker addresses and submit a hash, startindex, and length of entries.
	for i := 0; i < len(workerAddrs); i++ {
		startIndex := int64(i) * pages

		// This go routine submits values to the PODS not the ec2s.
		go func(index int) {

			resp, er := tryGet(workerAddrs[index], hash, startIndex, pages)

			if er != nil {
				log.Println(er)
			}
			result = append(result, Data{Hash: resp.Hash, Result: resp.Result})

		}(i)
	}
	return result
}

// tryGet attempts to send the hash string to the address every second, for t seconds. If no connection is made in that time, returns an error.
func tryGet(addr, hash string, index int64, length int64) (Data, error) {
	var er error
	var tmp Data
	var colliderPort string = "8080"

	// Submit request to colliders. Post request will be used because GET query and
	// POST Content-Type: application/x-www-form-urlencoded get parsed exactly the same by
	// myURL.PasreForm()
	colliderURL := fmt.Sprintf("http://%s:%s/", addr, colliderPort)
	stringIndex := string(index)
	stringLength := string(length)
	// contentType := "application/x-www-form-urlencoded"
	// content := fmt.Sprintf("hash=%s&index=%s&length=%s", hash, index, length)

	request, er := http.NewRequest("GET", colliderURL+"?hash="+hash+"&index="+stringIndex+"&length="+stringLength, nil)
	if er != nil {
		log.Println(er, "Backend server connection failed")
	}
	client := http.Client{}
	response, er := client.Do(request)
	json.NewDecoder(response.Body).Decode(&tmp)

	fmt.Println(tmp)

	return tmp, er
}

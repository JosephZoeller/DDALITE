package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"io"
	"log"
	"net/http"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

var collision []cityhashutil.HashCollision
/*
// Send hash to each worker's overlay ip address.
func sendToWorkersLegacy(hash string) string {
	responseChan := make(chan io.ReadCloser, len(overIps))
	var workerCount int64 = int64(len(overIps))  // How many slave ips have we registered?
	var pages int64 = dictionaryLength / workerCount // How big will the assigned tasks will be

	var tmp cityhashutil.HashCollision

	for i := 0; i < len(overIps); i++ {
		startIndex := int64(i) * pages

		// This go routine submits values to the PODS not the ec2s.
		go func(index int) {

			resp, er := tryGet(overIps[index], hash, startIndex, pages)
			if er != nil {
				log.Println(er)
			}
			responseChan <- resp.Body

		}(i)
	}

	for i := 0; i< len(overIps); i++ {
		firstResponseBody := <- responseChan
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
*/

func listenForWorker(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Received data from worker, passing it onto client...")
	collision := cityhashutil.HashOutParams{}
	err := json.NewDecoder(req.Body).Decode(&collision)
	if err != nil {
		fmt.Println(err)
	} else {
		clientMsg := cityhashutil.ResponseMessage{ Message: fmt.Sprint(collision.Hashed, " | ", collision.Unhashed)}
		post, _ := json.Marshal(clientMsg)
		http.Post("http://localhost:999/sdnc", "application/json", bytes.NewReader(post))
	}


}

func sendToWorkers(workSpec cityhashutil.ClientPost) {
	for i, addr := range overIps {
		if i < len(workSpec.Dictionaries) {
			work, _ := json.Marshal(cityhashutil.HashInParamsOnline{
				InputHashes: workSpec.InputHashes, 
				Dictionary: workSpec.Dictionaries[i], 
				Delimiter: workSpec.Delimiter, 
				Depth: workSpec.Depth,
			})

			msg := cityhashutil.ResponseMessage{}
			rsp, err := http.Post(addr, "application/json", bytes.NewReader(work))
			if err != nil {
				panic(err)
			} else {
				err = json.NewDecoder(rsp.Body).Decode(&msg)
				if err != nil {
					log.Println("Failed to decode server response - ", err)
				} else {
					log.Println(msg.Message)
				}
			}
		}
	}
}
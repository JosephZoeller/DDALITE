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

func listenForWorker(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Received data from worker...")
	collision := cityhashutil.HashOutParams{}
	err := json.NewDecoder(req.Body).Decode(&collision)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully decoded worker data. Passing it onto client..")
		clientMsg := cityhashutil.ResponseMessage{ Message: fmt.Sprint(collision.Hashed, " | ", collision.Unhashed)}
		post, _ := json.Marshal(clientMsg)
		postAddr := fmt.Sprintf("http://%s:8080/SDNCToClient", clientAddr)
		http.Post(postAddr, "application/json", bytes.NewReader(post))
	}

}

func sendToWorkers(workSpec cityhashutil.ClientPost) {
	for i, addr := range overIps {
		log.Println("Sending work to: ", addr)
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
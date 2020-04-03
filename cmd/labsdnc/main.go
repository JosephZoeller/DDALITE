package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

var (
	overIps          []string
)

func main() {
	overIps = []string {"http://localhost:8080"}

	http.HandleFunc("/client", listenForClient)
	http.HandleFunc("/worker", listenForWorker)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	go func() {
		err := http.ListenAndServe("localhost:666", nil)
		if err != nil {
			log.Println(err)
			signalChan <- os.Kill
		}
	}()

	<-signalChan
}

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

func listenForClient(rw http.ResponseWriter, req *http.Request) {
	workSpec := cityhashutil.ClientPost{}

	srvMsg := cityhashutil.ResponseMessage{}
	err := json.NewDecoder(req.Body).Decode(&workSpec)
	if err != nil {
		srvMsg.Message = "Failed to decode"
		json.NewEncoder(rw).Encode(srvMsg)
		log.Println("Failed to decode client Post.")
	} else {
		srvMsg.Message = "Successful decode"
		json.NewEncoder(rw).Encode(srvMsg)
	}

	sendToWorkers(workSpec)
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
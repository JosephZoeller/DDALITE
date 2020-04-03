package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

func main() {
	http.HandleFunc("/", listenForClient)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	
	go func() {
		err := http.ListenAndServe("localhost:666", nil)
		if err != nil {
			log.Println(err)
			signalChan <- os.Kill
		}
	}()

	<- signalChan
}

func listenForClient(rw http.ResponseWriter, req *http.Request) {
	client := cityhashutil.ClientPost{}

	srvMsg := cityhashutil.MessageToClient{}
	err := json.NewDecoder(req.Body).Decode(&client)
	if err != nil {
		srvMsg.Message = "Failed to decode"
		json.NewEncoder(rw).Encode(srvMsg)
		log.Println("Failed to decode client Post.")
	} else {
		srvMsg.Message = "Successful decode"
		json.NewEncoder(rw).Encode(srvMsg)
	}
}
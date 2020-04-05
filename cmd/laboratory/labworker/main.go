package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)
var SDNCAddr string

func main() {

	http.HandleFunc("/", ListenForSDNC)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func ListenForSDNC(rw http.ResponseWriter, req *http.Request) {
	log.Println("Request recieved from SDNC: ", req.RemoteAddr)
	SDNCAddr = strings.Split(req.RemoteAddr, ":")[0]
	work := cityhashutil.HashInParamsOnline{}
	msg := cityhashutil.ResponseMessage{}

	err := json.NewDecoder(req.Body).Decode(&work)
	if err != nil {
		msg.Message = "Failed to decode"
		json.NewEncoder(rw).Encode(msg)
		log.Println("Failed to decode request from SDNC.")
		return
	}
	collisionChan := make(chan cityhashutil.HashOutParams)
	go findCollisions(collisionChan)

	msg.Message = "Worker is seeking collisions..."
	json.NewEncoder(rw).Encode(msg)
}

func findCollisions(collissionChan chan cityhashutil.HashOutParams) {
	for i := 0; i < 5; i++ {
		postCollisions(cityhashutil.HashOutParams{Hashed: "1234", Unhashed: "TestUnhash", Err: ""})
		time.Sleep(time.Second * 10)
	}
}

func postCollisions(collision cityhashutil.HashOutParams) {
	post, err := json.Marshal(collision)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sending False Data...")
		postAddr := fmt.Sprintf("http://%s:666/WorkerToSDNC", SDNCAddr)
		http.Post(postAddr, "application/json", bytes.NewReader(post))
	}
}

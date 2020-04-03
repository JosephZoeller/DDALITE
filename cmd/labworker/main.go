package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

func main() {

	http.HandleFunc("/", ListenForSDNC)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func ListenForSDNC(w http.ResponseWriter, r *http.Request) {
	work := cityhashutil.HashInParamsOnline{}
	json.NewDecoder(r.Body).Decode(&work)

	collisionChan := make(chan cityhashutil.HashOutParams)
	go findCollisions(collisionChan)

	postCollisions(collisionChan)
}

func findCollisions(collissionChan chan cityhashutil.HashOutParams) {
	for {
		collissionChan <- cityhashutil.HashOutParams{Hashed: "1234", Unhashed: "TestUnhash", Err: ""}
		time.Sleep(time.Second * 10)
	}
}

func postCollisions(collissionChan chan cityhashutil.HashOutParams) {
	for {
		collision := <-collissionChan
		post, err := json.Marshal(collision)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Sending False Data...")
			http.Post("http://localhost:666/worker", "application/json", bytes.NewReader(post))
		}
	}
}

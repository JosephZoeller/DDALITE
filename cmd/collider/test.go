package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)


func AlgorithmTest(rw http.ResponseWriter, req *http.Request) {
	log.Println("Algorithm Test:", req.RemoteAddr)

	work := cityhashutil.ColliderSpecifications{
		InputHashes: []uint64{0},
		Dictionary:  []string{"This", "is", "a", "test"},
		Delimiter:   "_",
		Depth:       3,
	}
	findCollisions(work)

	log.Println("SuccessTest", req.RemoteAddr)

	work = cityhashutil.ColliderSpecifications{
		InputHashes: []uint64{85894109417755},
		Dictionary:  []string{"A", "p", "l", "e"},
		Delimiter:   "",
		Depth:       5,
	}

	findCollisions(work)
}

func postCollisionsTest() {
	for {
		post, err := json.Marshal(<-collisionChan)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Sending Local Data...")
			postAddr := "http://localhost:8081/LocalToClient"
			http.Post(postAddr, "application/json", bytes.NewReader(post))
		}
	}
}
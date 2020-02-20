package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/200106-uta-go/JKJP2/pkg/cityhashutil"
)

// Data is response data struct

// Resp responses to request in JSON format
func Resp(w http.ResponseWriter, r *http.Request) {

	// Parse request values
	hash := r.PostFormValue("hash")
	index, er := strconv.Atoi(r.PostFormValue("index"))
	if er != nil {
		fmt.Fprint(w, "")
		return
	}
	length, er := strconv.Atoi(r.PostFormValue("length"))
	if er != nil {
		fmt.Fprint(w, "")
		return
	}

	// Retrieve collision
	response := cityhashutil.HashCollision{
		InputHash: hash,
		Collision: findCollisionFile(hash, index, length),
	}

	// Generate http response in JSON format
	output, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusTeapot)
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(output))
}

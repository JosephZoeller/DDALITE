package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Data is response data struct
type Data struct {
	InputHash string `json:"hash"`
	Collision string `json:"collision"`
}

// Resp responses to request in JSON format
func Resp(w http.ResponseWriter, r *http.Request) {

	// Parse request values
	hash := r.PostFormValue("hash")
	index, er := strconv.Atoi(r.PostFormValue("index"))
	if (er != nil) {
		fmt.Fprint(w, "")
		return
	}
	length, er := strconv.Atoi(r.PostFormValue("length"))
	if (er != nil) {
		fmt.Fprint(w, "")
		return
	}

	// Retrieve collision
	response := Data {
		InputHash: hash,
		Collision: findCollision(hash, index, length),
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

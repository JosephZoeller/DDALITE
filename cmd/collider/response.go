package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

// Data is response data struct

// Resp responses to request in JSON format
func Resp(w http.ResponseWriter, r *http.Request) {
	// Parse request values
	errStr := ""
	hash := r.FormValue("hash")
	if hash == "" {
		errStr = "Parse error - hash is empty."
		hash = "1234"
	}
	index, er := strconv.Atoi(r.FormValue("index"))
	if er != nil {
		errStr = "Parse error - could not parse index."
		index = 0
	}
	length, er := strconv.Atoi(r.FormValue("length"))
	if er != nil {
		fmt.Fprint(w, "")
		errStr = "Parse error - could not parse length."
		length = 1
	}

	// Retrieve collision
	dictionary, er := getDictionary(dictionaryFilePath)
	if er != nil {
		errStr = "Dictionary Error - Could not get dictionary."
	}
	response := cityhashutil.HashCollision{
		InputHash: hash,
		Collision: findCollisionFile(dictionary, hash, index, length),
		Err:       errStr,
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

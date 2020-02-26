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
	hash := r.FormValue("hash")
	if hash == "" {
		fmt.Fprint(w, "Parse error - hash is empty.")
		return
	}
	index, er := strconv.Atoi(r.FormValue("index"))
	if er != nil {
		fmt.Fprint(w, "Parse error - could not parse index.")
		return
	}
	length, er := strconv.Atoi(r.FormValue("length"))
	if er != nil {
		fmt.Fprint(w, "")
		fmt.Fprint(w, "Parse error - could not parse index.")
		return
	}

	// Retrieve collision
	dictionary, er := getDictionary(dictionaryFilePath)
	if er != nil {
		fmt.Fprint(w, "Dictionary Error - Could not get dictionary.")
		return
	}
	response := cityhashutil.HashCollision{
		InputHash: hash,
		Collision: findCollisionFile(dictionary, hash, index, length),
	}

	// Generate http response in JSON format
	output, err := json.Marshal(response)
	if err != nil {
		fmt.Fprint(w, "Send Error - Could not marshal response.")
		w.WriteHeader(http.StatusTeapot)
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(output))
}

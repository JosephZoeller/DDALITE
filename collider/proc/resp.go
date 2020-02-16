package proc

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Data is response data struct
type Data struct {
	Hash   string `json:"hash"`
	Result string `json:"result"`
}

// Resp responses to request in JSON format
func Resp(w http.ResponseWriter, r *http.Request) {
	var response Data

	// Get user entry (hash).
	response.Hash = r.FormValue("hash")

	// Sleep for random amount of seconds to simulate processing time (maximum 5 seconds).
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(6)) * time.Second)

	// Set result value
	response.Result = "Revature Web Service KJ3 instance"

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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)
var SDNCAddr = ""
func main() {

	post, err := json.Marshal(cityhashutil.ClientPost{
		InputHashes:  []string{"85894109417755"},
		Dictionaries: [][]string{[]string{""}},
		Delimiter:    "",
		Depth:        4,
	})
	if err != nil {
		panic(err)
	}

	msg := cityhashutil.ResponseMessage{}
	rsp, err := http.Post(fmt.Sprintf("http://%s:666/ClientToSDNC", SDNCAddr), "application/json", bytes.NewReader(post))
	if err != nil {
		panic(err)
	} else {
		err = json.NewDecoder(rsp.Body).Decode(&msg)
		if err != nil {
			fmt.Println("Failed to decode server response")
		} else {
			fmt.Println(msg.Message)
		}
	}

	fmt.Println("DONE")
}

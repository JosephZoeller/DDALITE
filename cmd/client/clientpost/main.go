package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/JosephZoeller/DDALITE/pkg/cityhashutil"
)

var SDNCAddr = "192.168.1.20" // <INSERT SDNC ADDRESS HERE>
const workFilePath = "work_order.json"

func init() {
	flag.Parse()
}

func main() {
	switch flag.Arg(0) {
	case "seek":
		workFile, err := os.Open(workFilePath)
		if err != nil {
			fmt.Printf("%s Error - %s\n", workFilePath, err.Error())
			os.Exit(1)
		}
		seekReq(workFile)
	case "teardown":
		teardownReq()
	case "genjson":
		OutputJson()
	//case "workerTest":
	//	workerTestReq()
	default:
		fmt.Println(usage)
		return
	}

	fmt.Println("DONE")
}

func OutputJson() {
	post, err := json.Marshal(cityhashutil.ClientSpecifications{
		InputHashes:  []uint64{85894109417755},
		Dictionaries: [][]string{[]string{"A", "p", "l", "e"}},
		Delimiter:    "",
		Depth:        5,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(post))
}

func workerTestReq() {
	_, err := http.Post("http://localhost:8080/Test", "application/json", nil)
	if err != nil {
		panic(err)
	}
}

func seekReq(post io.Reader) {
	msg := cityhashutil.MessageResponse{}
	rsp, err := http.Post(fmt.Sprintf("http://%s:666/ClientToSDNC", SDNCAddr), "application/json", post)
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
}

func teardownReq() {
	msg := cityhashutil.MessageResponse{}
	rsp, err := http.Post(fmt.Sprintf("http://%s:666/teardown", SDNCAddr), "application/json", nil)
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
}

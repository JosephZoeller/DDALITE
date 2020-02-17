package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
	"time"
)

// listenForWorker awaits messages from workers. Upon recieving a message, the hash-collision pair is saved to a text file.
func listenForWorker() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var hash = r.FormValue("hash")
		var collision = r.FormValue("collision")
		exportCollision(hash, collision)
	})
	http.ListenAndServe(":8081", nil)
}

func sendToWorkers(hash string) {
	for i := 0; i < len(workerAddrs); i++ {
		go func(index int) {
			resp, er := tryGet(workerAddrs[index], hash, 5)
			if er == nil {
				resp.Write(os.Stdout)
			} else {
				log.Fatal("Timeout: failed to connect - ", er)
			}
		}(i)
	}
}

// tryGet attempts to send the hash string to the address every second, for t seconds. If no connection is made in that time, returns an error.
func tryGet(addr, hash string, t int) (*http.Response, error) {
	var er error
	var resp *http.Response

	i := 0
	for i < t {

		resp, er = http.Get(fmt.Sprintf("%s/?hash=%s", addr, hash))
		if er == nil {
			return resp, er
		}
		fmt.Println(er)
		time.Sleep(time.Second)
		i++
	}
	return nil, er
}

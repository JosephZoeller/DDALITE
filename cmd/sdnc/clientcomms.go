package main

import (
	"net/http"
)

// listenForClient awaits a query (curl request) from the client. Upon recieving a request, the hash is handed out to the worker addresses.
func listenForClient() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var hash = r.FormValue("hash")
		sendToWorkers(hash)
	})
	http.ListenAndServe(":8080", nil)
}

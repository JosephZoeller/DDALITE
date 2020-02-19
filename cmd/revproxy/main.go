package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	revproxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "18.216.49.100:30000", // Need to be changed to SDN.
	})

	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", revproxy))
}

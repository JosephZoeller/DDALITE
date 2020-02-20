package main

import (
	"net/http"
	"text/template"
)

// Template is a struct for HTML loading page
type Template struct {
	Hash      string
	Instances string
}

// Loading is HTML loading page template generator
func Loading(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(rw, "Error Parsing client request form.", http.StatusInternalServerError)
	}

	container := Template{
		Hash:      req.FormValue("hash"),
		Instances: req.FormValue("instances"),
	}
	template := template.Must(template.ParseFiles("cmd/sdnc/html/loading.html"))
	template.Execute(rw, container)
}

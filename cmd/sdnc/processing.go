package main

import (
	"net/http"
	"text/template"
)

// Tmpl is a struct for HTML loading page
type Tmpl struct {
	Hash      string
	Instances string
	Result    string
}

func processing(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(rw, "Error Parsing client request form.", http.StatusInternalServerError)
	}

	container := Tmpl{
		Hash:      req.FormValue("hash"),
		Instances: req.FormValue("instances"),
	}
	template := template.Must(template.ParseFiles("cmd/sdnc/html/loading.html"))
	template.Execute(rw, container)
}

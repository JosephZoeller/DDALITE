package main

import "net/http"

func loading(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(rw, "Error Parsing client request form.", http.StatusInternalServerError)
	}

	//hash := req.FormValue("hash")
	//instances := req.FormValue("instances")
	//http.FileServer(http.Dir("html/loading.html"))
	http.Redirect(rw, req, "https://google.com", http.StatusFound)
}

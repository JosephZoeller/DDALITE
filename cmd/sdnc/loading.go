package main

import "net/http"

func loading(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		http.Error(rw, "Error Parsing client request form.", http.StatusInternalServerError)
	}

	hash := req.FormValue("hash")
	instances := req.FormValue("instances")
	go http.FileServer(http.Dir("html/loading.html"))
	go http.Redirect(rw, req, "/result/?hash="+hash+"&instances="+instances, 418)
}

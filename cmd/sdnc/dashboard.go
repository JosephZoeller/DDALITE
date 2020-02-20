package main

import (
	"net/http"
	"text/template"

	"github.com/200106-uta-go/JKJP2/pkg/kubeutil"
)

type tmpl struct {
	Node    []kubeutil.Node
	Pod     []kubeutil.Pod
	Service []kubeutil.Service
}

func dashboard(rw http.ResponseWriter, req *http.Request) {
	container := tmpl{
		Node:    kubeutil.NodeInfo(),
		Pod:     kubeutil.PodInfo(),
		Service: kubeutil.ServiceInfo(),
	}
	template := template.Must(template.ParseFiles("cmd/sdnc/html/dash.html"))
	template.Execute(rw, container)
}

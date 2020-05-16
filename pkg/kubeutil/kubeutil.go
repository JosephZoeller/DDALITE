// Package kubeutil generic utilities for working with k8s
package kubeutil

import (
	"log"
	"strings"

	"github.com/JosephZoeller/DDALITE/pkg/shellutil"
)

// NodeInfo Retrieves Data on All nodes on cluster
func NodeInfo() []Node {
	var Nodes []Node

	output, err := shellutil.RunCommand("kubectl --kubeconfig=/home/pi/.kube/config get nodes -o wide", ".")
	log.Println(output)
	if err != nil {
		panic(err)
	}

	line := strings.Split(output, "\n")
	line = line[1:]

	for _, detail := range line {

		field := strings.Split(detail, "  ")

		var tmp []string
		for _, text := range field {
			if strings.TrimSpace(text) != "" {
				tmp = append(tmp, text)
			}
		}

		if len(tmp) > 0 {
			Nodes = append(Nodes, Node{
				Name:             strings.TrimSpace(tmp[0]),
				Status:           strings.TrimSpace(tmp[1]),
				Role:             strings.TrimSpace(tmp[2]),
				Age:              strings.TrimSpace(tmp[3]),
				Version:          strings.TrimSpace(tmp[4]),
				InternalIP:       strings.TrimSpace(tmp[5]),
				ExternalIP:       strings.TrimSpace(tmp[6]),
				OSImage:          strings.TrimSpace(tmp[7]),
				KernelVer:        strings.TrimSpace(tmp[8]),
				ContainerRunTime: strings.TrimSpace(tmp[9]),
			})
		}
	}
	return Nodes
}

// PodInfo retrieves data on all Pods on cluster
func PodInfo() []Pod {
	var Pods []Pod

	output, err := shellutil.RunCommand("kubectl --kubeconfig=/home/pi/.kube/config get pods -o wide", ".")
	log.Println(output)
	if err != nil {
		panic(err)
	}

	line := strings.Split(output, "\n")
	line = line[1:]

	for _, detail := range line {

		field := strings.Split(detail, "  ")

		var tmp []string
		for _, text := range field {
			if strings.TrimSpace(text) != "" {
				tmp = append(tmp, text)
			}
		}

		if len(tmp) > 0 {
			Pods = append(Pods, Pod{
				Name:          strings.TrimSpace(tmp[0]),
				Ready:         strings.TrimSpace(tmp[1]),
				Status:        strings.TrimSpace(tmp[2]),
				Restart:       strings.TrimSpace(tmp[3]),
				Age:           strings.TrimSpace(tmp[4]),
				IPaddr:        strings.TrimSpace(tmp[5]),
				Node:          strings.TrimSpace(tmp[6]),
				NominatedNode: strings.TrimSpace(tmp[7]),
				ReadinessGate: strings.TrimSpace(tmp[8]),
			})
		}
	}
	return Pods
}

// ServiceInfo retrieves data on all Services on cluster
func ServiceInfo() []Service {
	var Services []Service

	output, err := shellutil.RunCommand("kubectl --kubeconfig=/home/pi/.kube/config get services -o wide", ".")
	log.Println(output)
	if err != nil {
		panic(err)
	}

	line := strings.Split(output, "\n")
	line = line[1:]

	for _, detail := range line {

		field := strings.Split(detail, "  ")

		var tmp []string
		for _, text := range field {
			if strings.TrimSpace(text) != "" {
				tmp = append(tmp, text)
			}
		}

		if len(tmp) > 0 {
			Services = append(Services, Service{
				Name:       strings.TrimSpace(tmp[0]),
				Type:       strings.TrimSpace(tmp[1]),
				ClusterIP:  strings.TrimSpace(tmp[2]),
				ExternalIP: strings.TrimSpace(tmp[3]),
				Port:       strings.TrimSpace(tmp[4]),
				Age:        strings.TrimSpace(tmp[5]),
				Selector:   strings.TrimSpace(tmp[6]),
			})
		}
	}
	return Services
}

// Package kubeutil generic utilities for working with k8s
package kubeutil

import (
	"os/exec"
	"strings"
)

type Node struct {
	Name             string
	Status           string
	Role             string
	Age              string
	Version          string
	InternalIP       string
	ExternalIP       string
	OSImage          string
	KernelVer        string
	ContainerRunTime string
}

type Pod struct {
	Name          string
	Ready         string
	Status        string
	Restart       string
	Age           string
	IPaddr        string
	Node          string
	NominatedNode string
	ReadinessGate string
}

type Service struct {
	Name       string
	Type       string
	ClusterIP  string
	ExternalIP string
	Port       string
	Age        string
	Selector   string
}

func NodeInfo() []Node {
	var Nodes []Node

	output, _ := exec.Command("kubectl", "get", "nodes", "-o", "wide").Output()
	line := strings.Split(string(output), "\n")
	line = line[1:]

	for _, detail := range line {

		field := strings.Split(detail, " ")

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

func PodInfo() []Pod {
	var Pods []Pod

	output, _ := exec.Command("kubectl", "get", "pods", "-o", "wide").Output()
	line := strings.Split(string(output), "\n")
	line = line[1:]

	for _, detail := range line {

		field := strings.Split(detail, " ")

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

func ServiceInfo() []Service {
	var Services []Service

	output, _ := exec.Command("kubectl", "get", "services", "-o", "wide").Output()
	line := strings.Split(string(output), "\n")
	line = line[1:]

	for _, detail := range line {

		field := strings.Split(detail, " ")

		var tmp []string
		for _, text := range field {
			if strings.TrimSpace(text) != "" {
				tmp = append(tmp, text)
			}
		}

		if len(tmp) > 0 {
			Services = append(Services, Service{
				Name:       strings.TrimSpace(tmp[0]),
				Type:       strings.TrimSpace(tmp[0]),
				ClusterIP:  strings.TrimSpace(tmp[0]),
				ExternalIP: strings.TrimSpace(tmp[0]),
				Port:       strings.TrimSpace(tmp[0]),
				Age:        strings.TrimSpace(tmp[0]),
				Selector:   strings.TrimSpace(tmp[0]),
			})
		}
	}
	return Services
}

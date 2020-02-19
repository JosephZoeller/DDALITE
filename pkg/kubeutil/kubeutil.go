// Package kubeutil generic utilities for working with k8s
package kubeutil

import (
	"log"
	"os/exec"
)

// "k8s.io/client-go/pkg/apis/clientauthentication"
// "k8s.io/client-go/tools/clientcmd"

// SetUp will apply deployment to cluster.
func SetUp() {
	runKube := exec.Command("kubectl", "apply", "-f", "/kubernetes/deployment.yaml")
	err := runKube.Run()
	if err != nil {
		log.Panicf("Could not run kubectl, Error == %v", err)
	}
}

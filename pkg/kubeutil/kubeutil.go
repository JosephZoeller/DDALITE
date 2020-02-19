// Package kubeutil generic utilities for working with k8s
package kubeutil

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd"
)

// SetUp will apply deployment to cluster.
func SetUp(instanceCount int) []string {

	fmt.Println(clientcmd.RecommendedConfigDir)

	/*
		runKube := exec.Command("kubectl", "apply", "-f", "/kubernetes/deployment.yaml")
		err := runKube.Run()
		if err != nil {
			log.Panicf("Could not run kubectl, Error == %v", err)
		}
	*/
}

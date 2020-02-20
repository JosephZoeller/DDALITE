package kubeutil

import (
	"fmt"
	"log"
	"os/exec"
)

// TearDown will first drain and then delete all pods on slave nodes.
func TearDown() error {
	// kubectl scale deployment.v1.apps/collider-deployment --replicas 0
	podCount := "0"
	out, err := exec.Command("sudo", "kubectl", "scale", "deployment.v1.apps/collider-deployment", "--replicas", podCount).Output()
	if err != nil {
		return fmt.Errorf("Could not kubectl delete: Error==%v", err)
	}

	log.Printf("-Teardown Successful-\n%s", out)
	return nil
}

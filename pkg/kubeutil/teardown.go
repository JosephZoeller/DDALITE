package kubeutil

import (
	"fmt"
	"log"
	"os/exec"
)

// TearDown will first drain and then delete all pods on slave nodes.
func TearDown() error {
	out, err := exec.Command("kubectl", "delete", "deploy/collider-deployment", "svc/collider-service").Output()
	if err != nil {
		return fmt.Errorf("Could not kubectl delete: Error==%v", err)
	}

	log.Printf("-Teardown Successful-\n%s", out)
	return nil
}

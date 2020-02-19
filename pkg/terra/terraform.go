// Package terra will hand function related to manipulating terraform
// configuration files.
package terra

import (
	"fmt"
	"os"
	"os/exec"
)

// Provision will create all the slave nodes based on what client
// send in request. This function returns the ips of all the created slave nodes.
// May change to pass in master token.
func Provision(instanceCount string) {
	arg := fmt.Sprintf("echo $(jq '.user_count = %s' var.json) > var.json", instanceCount)
	cmd := exec.Command("bash", "-c", arg)
	cmd.Dir = "/home/ubuntu/terradir"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd = exec.Command("/bin/sh", "-c", "sudo terraform init")
	cmd.Dir = "/home/ubuntu/terradir"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd = exec.Command("/bin/sh", "-c", "sudo terraform apply -auto-approve")
	cmd.Dir = "/home/ubuntu/terradir"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

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
func Provision(instanceCount string) []string {
	// IF count=6 deploy 6 ec2s ELSE deploy 3 ec2s.
	if instanceCount == "6" {

		// cmd := exec.Command("bash", "-c", "echo hello > /tmp/aaa")

		cmd := exec.Command("bash", "-c", "echo $(jq '.user_count = \"6\"' var.json) > var.json")
		cmd.Dir = "/home/ubuntu/terradir"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else {

		// cmd := exec.Command("bash", "-c", "echo hello > /tmp/aaa")

		cmd := exec.Command("bash", "-c", "echo $(jq '.user_count = \"3\"' var.json) > var.json")
		cmd.Dir = "/home/ubuntu/terradir"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
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

	// cmd export instance_ips=$(terraform output instance_ips)
	cmd = exec.Command("/bin/sh", "-c", "export instance_ips=$(terraform output instance_ips)")
	cmd.Dir = "/home/ubuntu/terradir"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	// cmd print instance_ips
	fmt.Print(os.Getenv("instance_ips"))

	return os.Getenv("instance_ips")

}

// Package terra will hand function related to manipulating terraform
// configuration files.
package terra

import (
	"os/exec"
)

// Provision will create all the slave nodes based on what client
// send in request. This function returns the ips of all the created slave nodes.
// May change to pass in master token.
func Provision(instanceCount string) []string {
	if instanceCount >5 {
	

		exec.Command("echo","$(jq '.user_count = \"6\"' var.json)" ">" "var.json")

	}
	
	exec.Command("cd", "/home/ubunut/terradir")
	exec.Command("terraform", "init")
	exec.Command("terraform", "apply")

	exec.Command("export", "instance_ips=$(terraform output instance_ips)")

	return os.Getenv("instance_ips")


}

// Package terra will hand function related to manipulating terraform
// configuration files.
package terra

import (
	"fmt"

	"github.com/JosephZoeller/DDALITE/pkg/shellutil"
)

// Provision will create all the slave nodes based on what client
// send in request. This function returns the ips of all the created slave nodes.
// May change to pass in master token.
func Provision(userCount int) {
	str, err := shellutil.RunCommand("echo $(jq '.user_count = %d' var.json) > var.json", "/home/ubuntu/terradir")
	if err != nil {
		panic(str)
	}
	fmt.Println(str)
	str, err = shellutil.RunCommand("sudo terraform init", "/home/ubuntu/terradir")
	if err != nil {
		panic(str)
	}
	fmt.Println(str)
	str, err = shellutil.RunCommand("sudo terraform apply --auto-approve", "/home/ubuntu/terradir")
	if err != nil {
		panic(str)
	}
	fmt.Println(str)
}

// TearDown will remove slave instances.
func TearDown() {
	str, err := shellutil.RunCommand("sudo terraform destroy -auto-approve", "/home/ubuntu/terradir")
	if err != nil {
		panic(str)
	}
}

/*
func TearDown() {

	cmd := exec.Command("/bin/sh", "-c", "sudo terraform destroy -auto-approve")
	cmd.Dir = "/home/ubuntu/terradir"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func Provision(userCount int) {
	arg := fmt.Sprintf("echo $(jq '.user_count = %d' var.json) > var.json", userCount)
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
	cmd = exec.Command("/bin/sh", "-c", "sudo terraform apply --auto-approve")
	cmd.Dir = "/home/ubuntu/terradir"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
*/

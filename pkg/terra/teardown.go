package terra

import (
	"os"
	"os/exec"
)

// TearDown will remove slave instances.
func TearDown() {

	cmd := exec.Command("/bin/sh", "-c", "sudo terraform destroy -auto-approve")
	cmd.Dir = "/home/ubuntu/terradir"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

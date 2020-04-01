package shellutil

import (
	"os/exec"
)

func RunCommand(command, dir string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	return string(out), err
}
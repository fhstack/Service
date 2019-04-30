package utils

import (
	"bytes"
	"os/exec"
)

const (
	Shell string = "bash"
)

func Exec(args string) (string, string, error) {
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd := exec.Command(Shell, "-c", args)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

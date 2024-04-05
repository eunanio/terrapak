package helpers

import (
	"bytes"
	"errors"
	"os/exec"
)

type Cmd struct{
	Dir string
}

func (c *Cmd) Execute(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = c.Dir

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	if stderr.Len() != 0 {
		return stdout.String(), errors.New(stderr.String())
	}

	return stdout.String(), nil
}
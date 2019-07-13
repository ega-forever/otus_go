package internal

import (
	"os"
	"os/exec"
)

func RunProgram(command string, env []string) error {

	cmd := exec.Command(command)
	cmd.Env = env
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	return err
}

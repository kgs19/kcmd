package kcmd

import (
	"fmt"
)

// RunDockerVersion executes the 'docker version' command using the runCommandPrintOutput function.
// Will print the output of the command to the console (stdout).
// It returns an error if the command execution fails.
//
// Returns:
//
//	error: An error if the 'docker version' command fails to execute, otherwise nil.
func RunDockerVersion() error {
	cmdStr := "docker"
	cmdDir := ""

	args := []string{"version"}
	if err := runCommandPrintOutput(cmdStr, cmdDir, nil, args...); err != nil {
		return fmt.Errorf("failed to execute 'docker version' command: %w", err)
	}
	return nil
}

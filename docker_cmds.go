package kcmd

import (
	"fmt"
)

const dockerCmd = "docker"

// RunDockerCmd executes a Docker command in the specified directory.
//
// Parameters:
// - cmdDir: The directory where the Docker command will be executed.
// - envVars: A slice of environment variables to be set for the command execution.
// - args: A variadic list of arguments for the Docker command. Meaning every argument after the docker command itself
//
// Returns:
// - error: An error if the Docker command execution fails, otherwise nil.
//
// The function constructs the Docker command with the specified arguments and executes it
// in the provided directory. It uses the runCommandPrintOutput function to run the command
// and print its output.
func RunDockerCmd(cmdDir string, envVars []string, args ...string) error {
	if err := runCommandPrintOutput(dockerCmd, cmdDir, nil, args...); err != nil {
		return err
	}
	return nil
}

// RunDockerVersion executes the 'docker version' command using the runCommandPrintOutput function.
// Will print the output of the command to the console (stdout).
// It returns an error if the command execution fails.
//
// Returns:
//
//	error: An error if the 'docker version' command fails to execute, otherwise nil.
func RunDockerVersion() error {
	cmdDir := DefaultConfig.BaseCommandDir

	args := []string{"version"}
	if err := RunDockerCmd(cmdDir, nil, args...); err != nil {
		return fmt.Errorf("failed to execute 'docker version' command: %w", err)
	}
	return nil
}

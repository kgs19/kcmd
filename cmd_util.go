package kcmd

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

// execCommandWithEnv executes a shell command with specified environment variables
// and outputs the result to the provided writer. It logs the command details and handles errors.
//
// Parameters:
// - cmdStr: The command to be executed.
// - cmdDir: The directory in which to execute the command.
// - envVars: A slice of environment variables to set for the command. os level cmd added automatically.
// - output: An io.Writer where the command's standard output will be written.
// - args: Additional arguments to pass to the command.
//
// Returns:
// - error: An error if the command fails, otherwise nil.
//
// The function logs the command details, sets the specified environment variables,
// and pipes the command's output to the provided writer. If the command fails, it
// captures the standard error message and exit code, and returns a custom error.
func execCommandWithEnv(cmdStr string, cmdDir string, envVars []string, output io.Writer, args ...string) error {
	exitCode := 0
	var errb bytes.Buffer

	if DefaultConfig.LogCommand {
		// Log the command details
		logCmd(cmdStr, cmdDir, envVars, args...)
	}

	// Set up the command with the provided directory and arguments
	cmd := exec.Command(cmdStr, args...)
	cmd.Dir = cmdDir

	// Set the environment variables from envVars
	setCmdEnvVars(cmd, envVars)

	// pipe the commands output to the applications
	// standard output
	cmd.Stdout = output
	cmd.Stderr = &errb
	err := cmd.Run()

	if err != nil {
		exitCode = 1
		stdErrorMsg := errb.String()
		var exitError *exec.ExitError
		if errors.As(err, &exitError) { // errors.As() -> function allows you to extract a specific error type from the error chain
			exitCode = exitError.ExitCode() //try to get actual cmd exitCode
		}
		err := NewCommandError(stdErrorMsg, exitCode, cmdStr, args...)
		return err
	}
	return nil
}

// execShCommandEnvPrintOutput executes a shell command with specified environment variables
// and prints the output to the standard output. It uses the execCommandWithEnv function to
// handle the command execution and output.
//
// Parameters:
// - cmdStr: The command to be executed.
// - cmdDir: The directory in which to execute the command.
// - envVars: A slice of environment variables to set for the command.
// - args: Additional arguments to pass to the command.
//
// Returns:
// - error: An error if the command fails, otherwise nil.
//
// This function sets up the command with the provided directory and environment variables,
// and prints the command's output to the standard output. It is a convenience wrapper around
// execCommandWithEnv that directs the output to os.Stdout.
func execShCommandEnvPrintOutput(cmdStr string, cmdDir string, envVars []string, args ...string) error {
	//print the output to the standard output
	output := os.Stdout
	return execCommandWithEnv(cmdStr, cmdDir, envVars, output, args...)
}

func setCmdEnvVars(cmd *exec.Cmd, envVars []string) {
	cmd.Env = os.Environ()
	if envVars != nil && len(envVars) > 0 {
		for _, envVar := range envVars {
			cmd.Env = append(cmd.Env, envVar)
		}
	}
}

func cmdStrWithArgs(cmdStr string, args ...string) string {
	return cmdStr + " " + strings.Join(args, " ")
}

func logCmd(cmdStr string, cmdDir string, envVars []string, args ...string) {
	cmd := cmdStrWithArgs(cmdStr, args...)
	slog.Info("cmdDir: " + cmdDir + " - env: " + strings.Join(envVars, " "))
	slog.Info("Executing: " + cmd)
}

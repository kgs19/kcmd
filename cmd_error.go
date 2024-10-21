package kcmd

import "fmt"

// CommandError represents an error that occurred while executing a command.
// It includes the exit code, standard error message, and the command that was executed.
type CommandError struct {
	exitCode int
	errorMsg string
	command  string
	cmdDir   string
}

func (e *CommandError) Error() string {
	if e.cmdDir == "" {
		return fmt.Sprintf(
			"failed Command: \n%s\n"+
				"exit code: %d\n"+
				"error message: \n%s\n",
			e.command, e.exitCode, e.errorMsg)
	}
	return fmt.Sprintf(
		"failed Command: \n%s\n"+
			"exit code: %d\n"+
			"error message: \n%s\n"+
			"execution directory: %s\n",
		e.command, e.exitCode, e.errorMsg, e.cmdDir)
}

func NewCommandError(errorMsg string, exitCode int, cmdDir string, cmdStr string, args ...string) *CommandError {
	return &CommandError{
		exitCode: exitCode,
		errorMsg: errorMsg,
		cmdDir:   cmdDir,
		command:  cmdStrWithArgs(cmdStr, args...),
	}
}

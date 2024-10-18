package kcmd

import "fmt"

// CommandError represents an error that occurred while executing a command.
// It includes the exit code, standard error message, and the command that was executed.
type CommandError struct {
	exitCode    int
	stdErrorMsg string
	command     string
}

func (e *CommandError) Error() string {
	return fmt.Sprintf(
		"Failed Command: \n%s\n"+
			"Exit code: %d\n"+
			"Error message: \n%s\n",
		e.command, e.exitCode, e.stdErrorMsg,
	)
}

func NewCommandError(stdErrorMsg string, exitCode int, cmdStr string, args ...string) *CommandError {
	return &CommandError{
		exitCode:    exitCode,
		stdErrorMsg: stdErrorMsg,
		command:     cmdStrWithArgs(cmdStr, args...),
	}
}

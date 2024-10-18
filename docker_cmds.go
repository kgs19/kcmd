package kcmd

import (
	"fmt"
)

func RunDockerVersion() error {
	cmdStr := "docker"
	cmdDir := ""

	args := []string{"version"}
	if err := execShCommandEnvPrintOutput(cmdStr, cmdDir, nil, args...); err != nil {
		return fmt.Errorf("failed to execute 'docker version' command: %w", err)
	}
	return nil
}
package main

import (
	"github.com/kgs19/kcmd"
	"log"
)

func main() {
	// run helm version command with the use of RunHelmCmd function
	cmdDir := kcmd.DefaultConfig.BaseCommandDir
	//Specify the helm command arguments
	//meaning every argument after the helm command itself
	helmCmdArgs := []string{"version"}
	err := kcmd.RunHelmCmd(cmdDir, nil, helmCmdArgs...)
	if err != nil {
		log.Fatalf("Error executing 'helm version' command: %v", err)
	}
}

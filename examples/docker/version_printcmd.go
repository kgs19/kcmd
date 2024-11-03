package main

import (
	"github.com/kgs19/kcmd"
	"log"
)

// example usage of the kcmd package
func main() {

	// Use custom Config struct to enable command logging
	customKcmdConfig := kcmd.DefaultConfig
	customKcmdConfig.PrintCommandEnabled = true

	kcmd.SetConfig(customKcmdConfig)
	err := kcmd.RunDockerVersion()
	if err != nil {
		log.Fatalf("Error executing 'docker version' command: %v", err)
	}
}

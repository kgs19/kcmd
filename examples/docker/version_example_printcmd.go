package main

import (
	"github.com/kgs19/kcmd"
	"log"
)

// example usage of the kcmd package
func main() {
	// USe Config struct to enable command logging
	customKcmdConfig := kcmd.Config{PrintCommand: true}
	kcmd.SetConfig(customKcmdConfig)
	err := kcmd.RunDockerVersion()
	if err != nil {
		log.Fatalf("Error executing 'docker version' command: %v", err)
	}
}

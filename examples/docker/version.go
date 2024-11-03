package main

import (
	"github.com/kgs19/kcmd"
	"log"
)

// example usage of the kcmd package
func main() {
	err := kcmd.RunDockerVersion()
	if err != nil {
		log.Fatalf("Error executing 'docker version' command: %v", err)
	}
}

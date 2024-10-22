// Package kcmd provides a set of utility functions that act as wrappers around
// basic kubectl, Helm, and Docker commands. These utilities simplify the
// execution of common tasks related to Kubernetes, Helm, and Docker by
// providing a higher-level interface.
//
// The package includes functions for executing shell commands, managing
// environment variables, and handling command outputs, among other utilities.
//
// Example usage:
//
//	    import "github.com/kgs19/kcmd"
//
//	    // Execute a 'docker version' command
//	    err := kcmd.RunDockerVersion()
//	    if err != nil {
//		     log.Fatalf("Error executing 'docker version' command: %v", err)
//	    }
package kcmd

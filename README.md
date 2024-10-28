# kcmd
A Golang library that serves as a wrapper around basic kubectl, Helm, and Docker cli commands.

## ✅ Prerequisites
This library expects the following tools to be installed and available in your system's PATH:
 - `kubectl`: The Kubernetes command-line tool for managing Kubernetes clusters.
 - `helm`: The package manager for Kubernetes applications.
 - `docker`: The platform for building, shipping, and running containerized applications.

## 📦 Installation
To install the library, run the following command:
```bash
go get -u github.com/kgs19/kcmd
```

## 📜 Commands Overview
The `kcmd` library provides the following high-level commands:
- **Helm Commands:**
    - `RunHelmInstall`: Executes the `helm upgrade --install` command to install or upgrade a Helm chart.
    - `RunHelmCmd`: Executes any Helm command with specific args. Use for cases not already covered by the library.
- **Docker Commands:**
    - `RunDockerVersion`: Executes the `docker version` command to display Docker version information.

### By convention, 
all commands that are prefixed with `Run` will print the output of the commands.
all commands that are prefixed with `Get` will return the output of the commands.

## 💡 Examples
See [./examples](./examples) for example usage.
```go
package main

import (
	"github.com/kgs19/kcmd"
	"log"
)

func main() {
	err := kcmd.RunDockerVersion()
	if err != nil {
		log.Fatalf("Error executing 'docker version' command: %v", err)
	}
}
```

## ⚙️ Config struct
### Library Configuration
The library also provides a `Config` struct that can be used to configure the behavior of the library.

```go
// Config holds the configuration settings for the kcmd library.
type Config struct {
	PrintCommandEnabled bool   // Flag to enable or disable printing the command executed
	BaseCommandDir      string // Directory to use to execute the commands
}
```
 - `PrintCommandEnabled`: Flag to enable or disable printing the command executed. Default is `false`.
 - `BaseCommandDir`: Directory to use to execute the commands. Default is the directory of the executable file.

Example on how to use the `Config` struct:
```go
package main

import (
	"github.com/kgs19/kcmd"
	"log"
)

func main() {
	// Use custom Config struct to print the command executed and the cmd path
	customKcmdConfig := kcmd.Config{PrintCommandEnabled: true}
	kcmd.SetConfig(customKcmdConfig)
	err := kcmd.RunDockerVersion()
	if err != nil {
		log.Fatalf("Error executing 'docker version' command: %v", err)
	}
}

```

### Config struct - Environment Variables
The library also provides a way to configure the library using environment variables.
```bash
export KCMD_PRINT_COMMAND_ENABLED=true
export KCMD_BASE_COMMAND_DIR=/path/to/base/dir
```
# kcmd
A Golang library that serves as a wrapper around basic kubectl, Helm, and Docker commands.

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

## 💡 Examples

See [./examples](./examples) for example usage.

```go
func main() {
	err := kcmd.RunDockerVersion()
	if err != nil {
		log.Fatalf("Error executing 'docker version' command: %v", err)
	}
}
```

##  Config struct - Library Configuration
The library also provides a `Config` struct that can be used to configure the behavior of the library.
```go
// Config holds the configuration settings for the kcmd library.
type Config struct {
	PrintCommandEnabled bool   // Flag to enable or disable command logging
	BaseCommandDir      string // Directory to use by default to execute all commands
}
```
Example on how to use the `Config` struct:
```go
func main() {
	// USe Config struct to enable command logging
	customKcmdConfig := kcmd.Config{PrintCommandEnabled: true}
	kcmd.SetConfig(customKcmdConfig)
	err := kcmd.RunDockerVersion()
	if err != nil {
		log.Fatalf("Error executing 'docker version' command: %v", err)
	}
}
```

## Config struct - Environment Variables
The library also provides a way to configure the library using environment variables.
```bash
export KCMD_PRINT_COMMAND_ENABLED=true
export KCMD_BASE_COMMAND_DIR=/path/to/base/dir
```
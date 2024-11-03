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
    - `RunHelmCmd`: Executes any Helm command with specific args. Use for cases not already covered by the library.

### By convention, 
all commands that are prefixed with `Run` will print the output of the commands.
all commands that are prefixed with `Get` will return the output of the commands.

## 💡 Examples
[Here](examples/docker/version.go) is one of the simplest example of using the library:
More examples are available in the [./examples](./examples) directory.

## ⚙️ Config struct

### Library Configuration
The library also provides a `Config` struct that can be used to configure the behavior of the library.
Following are the fields of the `Config` struct:
 - `PrintCommandEnabled`: Flag to enable or disable printing the command executed. Default is `false`.
 - `BaseCommandDir`: Directory to use to execute the commands. Default is the directory of the executable file.
[Here](examples/docker/version_printcmd.go) is an example on how to use the `Config` struct:

### Config struct - Environment Variables
The library also provides a way to configure the library through environment variables.
```bash
export CMDX_PRINT_COMMAND_ENABLED=true
export KCMD_BASE_COMMAND_DIR=/path/to/base/dir
```


## TODO
- Add `Get` commands 
- Review Config struct - move logic to cmdx ??

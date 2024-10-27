# Examples

This directory contains example usage of the `kcmd` library. Below are the details of each example provided.

## Helm Commands

### Install Nginx Helm Chart

This example demonstrates how to use the `kcmd` library to install a [Nginx Helm chart](./helm/nginx-chart) .

#### File: [install_nginx.go](./helm/install_nginx.go)

```go
package main

import (
	"github.com/kgs19/kcmd"
	"log"
)

func main() {
	helmChart := kcmd.HelmChart{
		Namespace:         "default",
		ReleaseName:       "nginx",
		RelativeChartPath: "nginx-chart",
		ValuesFiles:       []string{"values.yaml", "values-dev.yaml"},
		SetValues:         map[string]string{"replicaCount": "2"},
		WaitFlag:          true,
	}
	err := kcmd.RunHelmInstall(helmChart)
	if err != nil {
		log.Fatalf("Error executing 'helm install' command: %v", err)
	}
}

```

#### Running the Example
- There are two ways to run the example:
  1. Run the example using the `go run` command.
  2. Build the example using the `go build` command and run the executable.

1. Run the example using the `go run` command.
   With this method, we need to set the following environment variables before running the example. 
```shell 
cd examples/helm/
export KCMD_PRINT_COMMAND_ENABLED=true # optional
export KCMD_BASE_COMMAND_DIR=$(pwd)
go run install_nginx.go
```

2. Build the example using the `go build` command and run the executable.
```shell
cd examples/helm/
go build install_nginx.go
./install_nginx
```

## Docker Commands

### Check Docker Version

This example demonstrates how to use the kcmd library to check the Docker version.
#### File: [version_example_printcmd.go](./docker/version_example_printcmd.go)
```go
package main

import (
	"github.com/kgs19/kcmd"
	"log"
)

// example usage of the kcmd package
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

#### Running the Example
Run the example using the `go run` command.
```shell
cd examples/docker/
go run version_example_printcmd.go

Executing cmd:
docker version

Client:
 Version:           27.2.0
... output truncated ... 
```

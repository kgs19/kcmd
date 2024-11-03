# Examples

This directory contains example usage of the `kcmd` library. Below are the details of each example provided.

## Examples - Helm Commands
### Example 1 - Install Nginx Helm Chart
- [Here](./helm/install_nginx_minimal.go) is an example to demonstrates how to use the `kcmd` library to install a nginx helm chart
- Running the Example
- There are two ways to run the example:
1. Run the example using the `go run` command.  
   With this method, we need to set the following environment variables before running the example. 
```shell 
cd examples/helm/
export CMDX_PRINT_COMMAND_ENABLED=true # optional
export KCMD_BASE_COMMAND_DIR=$(pwd)
go run install_nginx.go
```

2. Build the example using the `go build` command and run the executable.
```shell
cd examples/helm/
go build install_nginx.go
./install_nginx
```

### Example 2 - Install Nginx Helm Chart with additional arguments
- [This](./helm/install_nginx.go) is the same example as above where additional optional Helm arguments are provided. 
- 

### Example 3 -  Uninstall Nginx Helm Chart
- [This](./helm/uninstall_nginx.go) example demonstrates how to use the `kcmd` library to uninstall a nginx helm chart.
- Refer to the first example for instructions on how to run the example.

--- 
## Examples - Docker Commands

### Example 1 - Print Docker Version
- [This](./docker/version_printcmd.go) example demonstrates how to use the kcmd library to print the Docker version.
- Run the example using the `go run` command.
```shell
cd examples/docker/
go run version_example_printcmd.go

Executing cmd:
docker version

Client:
 Version:           27.2.0
... output truncated ... 
```

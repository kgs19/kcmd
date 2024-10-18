# kcmd
A Golang library that serves as a wrapper around basic kubectl, Helm, and Docker commands.

## Prerequisites
This library expects the following tools to be installed and available in your system's PATH:
 - `kubectl`: The Kubernetes command-line tool for managing Kubernetes clusters.
 - `helm`: The package manager for Kubernetes applications.
 - `docker`: The platform for building, shipping, and running containerized applications.

## Installation
To install the library, run the following command:
```bash
go get -u github.com/kgs19/kcmd
```
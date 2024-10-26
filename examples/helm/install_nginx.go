package main

import "github.com/kgs19/kcmd"

// Set the following environment variables before running the example
// export KCMD_PRINT_COMMAND_ENABLED=true
// export KCMD_BASE_COMMAND_DIR=$(pwd)

// example usage of the kcmd package
func main() {

	helmChart := kcmd.HelmChart{
		Namespace:         "default",
		ReleaseName:       "nginx",
		RelativeChartPath: "nginx-chart",
		ValuesFiles:       []string{"values.yaml", "values-dev.yaml"},
		SetValues:         map[string]string{"replicaCount": "2"},
		WaitFlag:          true,
	}
	kcmd.RunHelmInstall(helmChart)
}
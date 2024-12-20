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
		OptionalHelmArgs:  map[string]string{"--timeout": "120s"},
		OptionalHelmFlags: []string{"--wait"},
	}
	err := kcmd.RunHelmInstall(helmChart)
	if err != nil {
		log.Fatalf("Error executing 'helm install' command: %v", err)
	}
}

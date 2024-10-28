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
	}

	err := kcmd.RunHelmInstall(helmChart)
	if err != nil {
		log.Fatalf("Error executing 'helm install' command: %v", err)
	}
}

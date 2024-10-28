package main

import (
	"github.com/kgs19/kcmd"
	"log"
)

func main() {

	helmChart := kcmd.HelmChart{
		Namespace:   "default",
		ReleaseName: "nginx",
	}
	err := kcmd.RunHelmUninstall(helmChart)
	if err != nil {
		log.Fatalf("Error executing 'helm install' command: %v", err)
	}
}

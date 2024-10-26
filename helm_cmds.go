package kcmd

import (
	"fmt"
)

// HelmChart represents the configuration for a Helm chart installation.
// It includes the release name, chart path, values files, set values, and a wait flag.
type HelmChart struct {
	Namespace         string            // Namespace is the Kubernetes namespace where the Helm chart will be installed.
	ReleaseName       string            // ReleaseName is the Helm release name.
	RelativeChartPath string            // RelativeChartPath is the relative path to the chart directory. It is relative to the BaseCommandDir set in the Config.
	ValuesFiles       []string          // ValuesFiles is a list of values files to be applied via the --values flag (e.g., values.yaml, values-dev.yaml).
	SetValues         map[string]string // SetValues is a map of key-value pairs to be applied via the --set flag.
	WaitFlag          bool              // WaitFlag indicates whether to set the --wait flag for the Helm install command.
}

func RunHelmInstall(helmChart HelmChart) error {
	relativePath := helmChart.RelativeChartPath
	releaseName := helmChart.ReleaseName
	namespace := helmChart.Namespace

	cmdStr := "helm"
	cmdDir := DefaultConfig.BaseCommandDir + "/" + relativePath

	args := []string{
		"upgrade", "--install", "--force", "--create-namespace",
		"--namespace", namespace, releaseName, ".",
	}

	for _, valueFile := range helmChart.ValuesFiles {
		args = append(args, "--values", valueFile)
	}

	//Remember the "--set" setting take precedence over the values files
	for key, value := range helmChart.SetValues {
		args = append(args, "--set", fmt.Sprintf("%s=%s", key, value))
	}

	if helmChart.WaitFlag {
		args = append(args, "--wait")
	}

	if err := runCommandPrintOutput(cmdStr, cmdDir, nil, args...); err != nil {
		return fmt.Errorf("failed to execute helm install with custom values, path=%s,: \n%w", relativePath, err)
	}
	return nil
}

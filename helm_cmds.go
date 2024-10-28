package kcmd

import (
	"fmt"
)

const helmCmd = "helm"

// HelmChart represents the configuration for a Helm chart installation.
// It includes the release name, chart path, values files, set values, and a wait flag.
type HelmChart struct {
	Namespace         string            // Namespace is the Kubernetes namespace where the Helm chart will be installed.
	ReleaseName       string            // ReleaseName is the Helm release name.
	RelativeChartPath string            // RelativeChartPath is the relative path to the chart directory. It is relative to the BaseCommandDir set in the Config.
	ValuesFiles       []string          // ValuesFiles is a list of values files to be applied via the --values flag (e.g., values.yaml, values-dev.yaml).
	SetValues         map[string]string // SetValues is a map of key-value pairs to be applied via the --set flag.
	OptionalHelmArgs  map[string]string // helmArgs is a map of key-value pairs representing additional arguments to be passed to the Helm command.
	OptionalHelmFlags []string          // helmFlags is a list of additional flags to be passed to the Helm command.
}

// RunHelmCmd executes a Helm command in the specified directory.
//
// Parameters:
// - cmdDir: The directory where the Helm command will be executed.
// - envVars: A slice of environment variables to be set for the command execution.
// - args: A variadic list of arguments for the Helm command. Meaning every argument after the helm command itself
//
// Returns:
// - error: An error if the Helm command execution fails, otherwise nil.
//
// The function constructs the Helm command with the specified arguments and executes it
// in the provided directory. It uses the runCommandPrintOutput function to run the command
// and print its output.
func RunHelmCmd(cmdDir string, envVars []string, args ...string) error {
	if err := runCommandPrintOutput(helmCmd, cmdDir, nil, args...); err != nil {
		return err
	}
	return nil
}

// RunHelmInstall installs or upgrades a Helm chart using the provided HelmChart configuration.
//
// Parameters:
// - helmChart: A HelmChart struct containing the configuration for the Helm chart installation.
//
// Returns:
// - error: An error if the Helm install command fails, otherwise nil.
//
// The function constructs the Helm install command with the specified release name, namespace,
// values files, set values, and wait flag. It then executes the command in the directory specified
// by the BaseCommandDir in the Config.
func RunHelmInstall(helmChart HelmChart) error {
	relativePath := helmChart.RelativeChartPath
	releaseName := helmChart.ReleaseName
	namespace := helmChart.Namespace

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

	for key, value := range helmChart.OptionalHelmArgs {
		args = append(args, key, value)
	}

	args = append(args, helmChart.OptionalHelmFlags...)

	if err := RunHelmCmd(cmdDir, nil, args...); err != nil {
		return fmt.Errorf("failed to execute helm install, path=%s,: \n%w", relativePath, err)
	}
	return nil
}

func RunHelmUninstall(helmChart HelmChart) error {
	relativePath := helmChart.RelativeChartPath
	releaseName := helmChart.ReleaseName
	namespace := helmChart.Namespace
	cmdDir := DefaultConfig.BaseCommandDir + "/" + relativePath

	//Examples:
	//helm -n=mnc-synca-gnss uninstall timescaledb
	args := []string{"uninstall", "--namespace", namespace, releaseName,
		// We are using --ignore-not-found since we do not want the ctl command to fail if one release is not found
		// TODO improvement for the future
		// avoid using the --ignore-not-found flag and programmatically check if the release exists
		// if the release does not exist, then log a info message
		"--ignore-not-found"}

	for key, value := range helmChart.OptionalHelmArgs {
		args = append(args, key, value)
	}

	args = append(args, helmChart.OptionalHelmFlags...)

	if err := RunHelmCmd(cmdDir, nil, args...); err != nil {
		return fmt.Errorf("failed to execute helm uninstall, path=%s,: \n%w", relativePath, err)
	}
	return nil
}

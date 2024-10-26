package kcmd

import (
	"testing"
)

func TestRunHelmInstall(t *testing.T) {
	// Mock HelmChart data
	helmChart := HelmChart{
		Namespace:         "default",
		ReleaseName:       "my-release",
		RelativeChartPath: "charts/my-chart",
		ValuesFiles:       []string{"values.yaml", "values-dev.yaml"},
		SetValues:         map[string]string{"key1": "value1"},
		WaitFlag:          true,
	}

	// Mock DefaultConfig
	DefaultConfig = Config{
		BaseCommandDir: "/path/to/base/command/dir",
	}

	// Mock runCommandPrintOutput function
	runCommandPrintOutput = func(cmdStr, cmdDir string, env []string, args ...string) error {
		// Verify the command and arguments
		expectedCmdStr := "helm"
		expectedCmdDir := "/path/to/base/command/dir/charts/my-chart"
		expectedArgs := []string{
			"upgrade", "--install", "--force", "--create-namespace",
			"--namespace", "default", "my-release", ".",
			"--values", "values.yaml",
			"--values", "values-dev.yaml",
			"--set", "key1=value1",
			"--wait",
		}

		if cmdStr != expectedCmdStr {
			t.Errorf("expected command %s, got %s", expectedCmdStr, cmdStr)
		}
		if cmdDir != expectedCmdDir {
			t.Errorf("expected command directory %s, got %s", expectedCmdDir, cmdDir)
		}
		for i, arg := range expectedArgs {
			if args[i] != arg {
				t.Errorf("expected argument %s, got %s", arg, args[i])
			}
		}
		return nil
	}

	// Run the function
	err := RunHelmInstall(helmChart)
	if err != nil {
		t.Errorf("RunHelmInstall() error = %v", err)
	}
}

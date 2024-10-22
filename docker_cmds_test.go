package kcmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunDockerVersion(t *testing.T) {
	// Save the original function to restore it later
	originalRunCommandPrintOutput := runCommandPrintOutput

	// Override the runCommandPrintOutput variable
	var output bytes.Buffer
	runCommandPrintOutput = func(cmdStr string, cmdDir string, envVars []string, args ...string) error {
		return runCommandWithEnv(cmdStr, cmdDir, envVars, &output, args...)
	}

	// Ensure the original function is restored after the test
	defer func() {
		runCommandPrintOutput = originalRunCommandPrintOutput
	}()

	tests := []struct {
		name               string
		config             Config
		expectedOutputKeys []string
		expectErr          bool
	}{
		{
			name:   "Default configuration",
			config: Config{},
			expectedOutputKeys: []string{
				"Client:",
				"Version:",
				"Go version:",
				"GitCommit:",
			},
			expectErr: false,
		},
		{
			name:   "Custom configuration with PrintCommandEnabled",
			config: Config{PrintCommandEnabled: true},
			expectedOutputKeys: []string{
				"Client:",
				"Version:",
				"Go version:",
				"GitCommit:",
				"docker version",
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetConfig(tt.config)
			err := RunDockerVersion()
			if (err != nil) != tt.expectErr {
				t.Errorf("RunDockerVersion() failed, expected error: %v, got: %v", tt.expectErr, err)
			}

			outputStr := output.String()
			//fmt.Printf("Output: %s\n", outputStr)
			// Check the captured output
			for _, key := range tt.expectedOutputKeys {
				if !strings.Contains(outputStr, key) {
					t.Errorf("Expected output to contain: %s, but it was not found", key)
				}
			}

		})
	}
}

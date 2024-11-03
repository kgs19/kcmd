package kcmd

import (
	"github.com/kgs19/cmdx"
	"strings"
	"testing"
)

func TestRunDockerVersion(t *testing.T) {
	// Save the original function to restore it later
	originalRunDockerCmd := RunDockerCmd

	// Ensure the original function is restored after the test
	defer func() {
		RunDockerCmd = originalRunDockerCmd
	}()

	tests := []struct {
		name               string
		config             cmdx.Config
		expectedOutputKeys []string
		expectErr          bool
	}{
		{
			name:   "Default configuration",
			config: cmdx.Config{},
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
			config: cmdx.Config{PrintCommandEnabled: true},
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
			cmdx.SetConfig(tt.config)
			// Mock the RunDockerCmd function to use RunCommandReturnOutput
			RunDockerCmd = func(cmdDir string, envVars []string, args ...string) error {
				output, err := cmdx.RunCommandReturnOutput(dockerCmd, cmdDir, nil, args...)
				if err != nil {
					return err
				}
				// Check the output for expected keys
				expectedOutputKeys := tt.expectedOutputKeys
				for _, key := range expectedOutputKeys {
					if !strings.Contains(output, key) {
						t.Errorf("Expected output to contain: %s, but it was not found", key)
					}
				}
				return nil
			}

			err := RunDockerVersion()
			if (err != nil) != tt.expectErr {
				t.Errorf("RunDockerVersion() failed, expected error: %v, got: %v", tt.expectErr, err)
			}
		})

	}
}

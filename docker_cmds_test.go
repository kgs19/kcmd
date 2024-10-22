package kcmd

import "testing"

func TestRunDockerVersion(t *testing.T) {
	tests := []struct {
		name      string
		config    Config
		expectErr bool
	}{
		{
			name:      "Default configuration",
			config:    Config{},
			expectErr: false,
		},
		{
			name:      "Custom configuration with PrintCommandEnabled enabled",
			config:    Config{PrintCommandEnabled: true},
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
		})
	}
}

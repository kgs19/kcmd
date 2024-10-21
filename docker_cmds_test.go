package kcmd

import "testing"

func TestRunDockerVersion(t *testing.T) {
	err := RunDockerVersion()
	if err != nil {
		t.Errorf("RunDockerVersion() failed, expected nil, got %v", err)
	}
}

package kcmd

// Config holds the configuration settings for the kcmd library.
type Config struct {
	PrintCommand bool // Flag to enable or disable command logging
}

// DefaultConfig provides default settings for the library.
var DefaultConfig = Config{
	PrintCommand: false, // Not print the command by default
}

// SetConfig allows users to set custom configuration options.
func SetConfig(cfg Config) {
	DefaultConfig = cfg
}

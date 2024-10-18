package kcmd

// Config holds the configuration settings for the kcmd library.
type Config struct {
	LogCommand bool // Flag to enable or disable command logging
}

// DefaultConfig provides default settings for the library.
var DefaultConfig = Config{
	LogCommand: false, // Default logging is disabled
}

// SetConfig allows users to set custom configuration options.
func SetConfig(cfg Config) {
	DefaultConfig = cfg
}

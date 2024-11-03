package kcmd

import (
	"github.com/kgs19/cmdx"
	"os"
	"path/filepath"
	"strconv"
)

// Config holds the configuration settings for the kcmd library.
type Config struct {
	cmdx.Config
	BaseCommandDir string // Directory to use to execute the commands
}

// DefaultConfig provides default settings for the library.
var DefaultConfig = Config{
	BaseCommandDir: getDefaultBaseCommandDir(),
}

// SetConfig allows users to set custom configuration options.
func SetConfig(cfg Config) {
	DefaultConfig = cfg
	cmdx.SetConfig(cfg.Config)
}

// getDefaultBaseCommandDir returns the default base command directory.
// It first checks the BASE_COMMAND_DIR environment variable.
// If not set, it defaults to the directory of the executable.
//
// Returns:
//
//	string: The base command directory.
func getDefaultBaseCommandDir() string {
	// Check if the KCMD_BASE_COMMAND_DIR environment variable is set
	baseCommandDir := getEnv("KCMD_BASE_COMMAND_DIR", ".")

	// If the environment variable is not set, use the directory of the executable
	if baseCommandDir == "." {
		ex, err := os.Executable()
		if err != nil {
			panic(err) // Panic if the executable path cannot be determined
		}
		executablePath := filepath.Dir(ex)
		baseCommandDir = executablePath
	}

	return baseCommandDir
}

// getEnv retrieves the value of the environment variable named by the key.
// If the variable is present in the environment, the function returns its value.
// Otherwise, it returns the specified default value.
//
// Parameters:
//   - key: The name of the environment variable to look up.
//   - defaultVal: The value to return if the environment variable is not set.
//
// Returns:
//
//	string: The value of the environment variable or the default value if the variable is not set.
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// Same as getEnv but returns a boolean value.
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

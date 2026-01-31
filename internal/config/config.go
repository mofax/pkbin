package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tailscale/hujson"
)

// Config represents the pkbin.jsonc configuration structure.
type Config struct {
	Scripts map[string]string `json:"scripts"`
}

// LoadConfig reads and parses pkbin.jsonc from the current working directory.
// Returns an error if the file is not found, invalid JSONC, or cannot be parsed.
func LoadConfig() (*Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current working directory: %w", err)
	}

	configPath := filepath.Join(cwd, "pkbin.jsonc")
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("pkbin.jsonc not found in current directory")
		}
		return nil, fmt.Errorf("failed to read pkbin.jsonc: %w", err)
	}

	// Parse JSONC format using hujson
	standardized, err := hujson.Standardize(data)
	if err != nil {
		// hujson errors already include line/column information
		return nil, fmt.Errorf("invalid JSONC in pkbin.jsonc: %w", err)
	}

	var config Config
	if err := json.Unmarshal(standardized, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal pkbin.jsonc: %w", err)
	}

	// Ensure Scripts map is initialized even if missing from JSON
	if config.Scripts == nil {
		config.Scripts = make(map[string]string)
	}

	return &config, nil
}

// FindScript looks up a script by name in the config.
// Returns the script command and an error if the script is not found.
func (c *Config) FindScript(name string) (string, error) {
	script, exists := c.Scripts[name]
	if !exists {
		return "", fmt.Errorf("script '%s' not found", name)
	}
	return script, nil
}

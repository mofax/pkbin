package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig_ValidJSONC(t *testing.T) {
	// Create a temporary directory
	tmpDir := t.TempDir()
	
	// Change to temp directory
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(oldDir)
	
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Create a valid pkbin.jsonc file with comments
	configContent := `{
  // This is a comment
  "scripts": {
    "build": "make build",
    "test": "pytest",
    "lint": "ruff check ."
  }
}`
	
	configPath := filepath.Join(tmpDir, "pkbin.jsonc")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() failed: %v", err)
	}

	if config == nil {
		t.Fatal("LoadConfig() returned nil config")
	}

	if len(config.Scripts) != 3 {
		t.Errorf("Expected 3 scripts, got %d", len(config.Scripts))
	}

	if config.Scripts["build"] != "make build" {
		t.Errorf("Expected build script 'make build', got '%s'", config.Scripts["build"])
	}

	if config.Scripts["test"] != "pytest" {
		t.Errorf("Expected test script 'pytest', got '%s'", config.Scripts["test"])
	}

	if config.Scripts["lint"] != "ruff check ." {
		t.Errorf("Expected lint script 'ruff check .', got '%s'", config.Scripts["lint"])
	}
}

func TestLoadConfig_InvalidJSONC(t *testing.T) {
	tmpDir := t.TempDir()
	
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(oldDir)
	
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Create an invalid JSONC file (malformed JSON)
	configContent := `{
  "scripts": {
    "build": "make build",
    // Missing closing brace
}`
	
	configPath := filepath.Join(tmpDir, "pkbin.jsonc")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	config, err := LoadConfig()
	if err == nil {
		t.Error("LoadConfig() should have failed with invalid JSONC")
	}
	if config != nil {
		t.Error("LoadConfig() should return nil config on error")
	}
}

func TestLoadConfig_MissingFile(t *testing.T) {
	tmpDir := t.TempDir()
	
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(oldDir)
	
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Don't create pkbin.jsonc file
	config, err := LoadConfig()
	if err == nil {
		t.Error("LoadConfig() should have failed when file is missing")
	}
	if config != nil {
		t.Error("LoadConfig() should return nil config on error")
	}

	// Check error message
	if err.Error() != "pkbin.jsonc not found in current directory" {
		t.Errorf("Expected 'pkbin.jsonc not found in current directory', got '%s'", err.Error())
	}
}

func TestLoadConfig_EmptyScripts(t *testing.T) {
	tmpDir := t.TempDir()
	
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(oldDir)
	
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Create config with empty scripts object
	configContent := `{
  "scripts": {}
}`
	
	configPath := filepath.Join(tmpDir, "pkbin.jsonc")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() failed: %v", err)
	}

	if config == nil {
		t.Fatal("LoadConfig() returned nil config")
	}

	if config.Scripts == nil {
		t.Error("Scripts map should be initialized even if empty")
	}

	if len(config.Scripts) != 0 {
		t.Errorf("Expected 0 scripts, got %d", len(config.Scripts))
	}
}

func TestLoadConfig_MissingScriptsField(t *testing.T) {
	tmpDir := t.TempDir()
	
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(oldDir)
	
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	// Create config without scripts field
	configContent := `{
  "other": "field"
}`
	
	configPath := filepath.Join(tmpDir, "pkbin.jsonc")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() failed: %v", err)
	}

	if config == nil {
		t.Fatal("LoadConfig() returned nil config")
	}

	// Scripts should be initialized as empty map
	if config.Scripts == nil {
		t.Error("Scripts map should be initialized even if missing from JSON")
	}
}

func TestConfig_FindScript(t *testing.T) {
	tmpDir := t.TempDir()
	
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(oldDir)
	
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp directory: %v", err)
	}

	configContent := `{
  "scripts": {
    "build": "make build",
    "test": "pytest"
  }
}`
	
	configPath := filepath.Join(tmpDir, "pkbin.jsonc")
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() failed: %v", err)
	}

	// Test finding existing script
	script, err := config.FindScript("build")
	if err != nil {
		t.Fatalf("FindScript('build') failed: %v", err)
	}
	if script != "make build" {
		t.Errorf("Expected 'make build', got '%s'", script)
	}

	// Test finding non-existent script
	_, err = config.FindScript("nonexistent")
	if err == nil {
		t.Error("FindScript('nonexistent') should have failed")
	}
	if err.Error() != "script 'nonexistent' not found" {
		t.Errorf("Expected 'script 'nonexistent' not found', got '%s'", err.Error())
	}
}

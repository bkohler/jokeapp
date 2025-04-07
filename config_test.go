package main

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

// TestGetConfigPath verifies the config file path is correctly generated in the user's home directory
func TestGetConfigPath(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Fatalf("Failed to get current user: %v", err)
	}
	expected := filepath.Join(usr.HomeDir, ".jokeapp.yaml")
	actual := getConfigFilePath()
	if actual != expected {
		t.Errorf("Expected config path %s, got %s", expected, actual)
	}
}

// TestReadAPIKeyFromConfig verifies reading an existing API key from the config file
func TestReadAPIKeyFromConfig(t *testing.T) {
	// Setup: create a temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".jokeapp.yaml")
	content := "deepseek_api_key: testkey123\n"
	if err := os.WriteFile(configPath, []byte(content), 0600); err != nil {
		t.Fatalf("Failed to write temp config: %v", err)
	}

	// Override home directory lookup
	origGetUser := getUserHomeDir
	getUserHomeDir = func() string { return tmpDir }
	defer func() { getUserHomeDir = origGetUser }()

	key, err := loadAPIKey()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if key != "testkey123" {
		t.Errorf("Expected API key 'testkey123', got '%s'", key)
	}
}

// Placeholder for prompt test
func TestPromptForAPIKeyIfMissing(t *testing.T) {
	t.Skip("Prompt test not implemented yet")
}

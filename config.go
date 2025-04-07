package main

import (
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/spf13/viper"
)

// getUserHomeDir returns the current user's home directory
var getUserHomeDir = func() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error fetching user home directory:", err)
		return ""
	}
	return usr.HomeDir
}

// getConfigFilePath returns the full path to the .jokeapp.yaml config file
func getConfigFilePath() string {
	homeDir := getUserHomeDir()
	return filepath.Join(homeDir, ".jokeapp.yaml")
}

// loadAPIKey attempts to load the Deepseek API key from the config file
func loadAPIKey() (string, error) {
	configPath := getConfigFilePath()
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}
	key := viper.GetString("deepseek_api_key")
	return key, nil
}

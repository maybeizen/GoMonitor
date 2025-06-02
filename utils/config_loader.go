package utils

import (
	"encoding/json"
	"fmt"
	"monitor/models"
	"os"
)

func LoadConfig(path string) (*models.Config, error) {
	// Default config
	config := &models.Config{
		MonitorInterval:   3,
		OutputPath:        "data/data.json",
		LogLevel:          "info",
		IncludeNetworks:   true,
		IncludeProcesses:  true,
		MaxProcessCount:   1000,
		EnableCompression: false,
	}

	// Try to read config file
	data, err := os.ReadFile(path)
	if err != nil {
		// If file doesn't exist, create it with default values
		if os.IsNotExist(err) {
			fmt.Printf("Config file not found at %s, creating with defaults...\n", path)
			return config, SaveConfig(config, path)
		}
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Parse JSON
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}

func SaveConfig(config *models.Config, path string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	return os.WriteFile(path, data, 0644)
}
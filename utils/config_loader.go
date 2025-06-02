package utils

import (
	"encoding/json"
	"fmt"
	"monitor/models"
	"os"
)

func LoadConfig(path string) (*models.Config, error) {
	config := &models.Config{
		MonitorInterval:   3,
		Outputs: []models.OutputConfig{
			{
				Type:     models.FileOutput,
				FilePath: "data/data.json",
			},
		},
		LogLevel:          "info",
		IncludeNetworks:   false,
		IncludeProcesses:  true,
		MaxProcessCount:   1000,
		EnableCompression: false,
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Config file not found at %s, creating with defaults...\n", path)
			return config, SaveConfig(config, path)
		}
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
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

func validateConfig(config *models.Config) error {
	if config.MonitorInterval < 1 {
		return fmt.Errorf("monitor_interval must be at least 1 second")
	}

	if len(config.Outputs) == 0 {
		return fmt.Errorf("at least one output must be configured")
	}

	for i, output := range config.Outputs {
		switch output.Type {
		case models.FileOutput:
			if output.FilePath == "" {
				return fmt.Errorf("output %d: file_path must be specified for file output", i)
			}
		case models.APIOutput:
			if output.APIURL == "" {
				return fmt.Errorf("output %d: api_url must be specified for API output", i)
			}
		default:
			return fmt.Errorf("output %d: unknown output type: %s", i, output.Type)
		}
	}

	return nil
}

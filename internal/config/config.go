package config

import (
	"encoding/json"
	"os"
)

// Config holds the configuration parameters from config.json
type Config struct {
	Port string `json:"port"`
}

// LoadConfig loads the configuration from the provided file path
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

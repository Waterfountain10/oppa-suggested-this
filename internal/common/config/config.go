package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// LoadConfig loads a YAML configuration file into the provided struct
func LoadConfig(path string, cfg interface{}) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return fmt.Errorf("error reading config file: %w", err)
    }

    if err := yaml.Unmarshal(data, cfg); err != nil {
        return fmt.Errorf("error parsing config file: %w", err)
    }

    return nil
}

package internal

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration for the monitoring tool.
type Config struct {
	URLs     []string `yaml:"urls"`
	Interval int      `yaml:"interval"`
	Timeout  int      `yaml:"timeout"`
	Logfile  string   `yaml:"logfile"`
}

// ReadConfig reads the configuration from the specified YAML file.
func ReadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

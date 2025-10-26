package helper

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MainDir string `yaml:"mainDir"`
}

func SaveConfig(cfg Config) error {
	configPath := filepath.Join(os.Getenv("HOME"), ".miyako-git-config.yaml")
	data, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}

func LoadConfig() (*Config, error) {
	configPath := filepath.Join(os.Getenv("HOME"), ".miyako-git-config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

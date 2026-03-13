package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Default []string `json:"default"`
}

func configPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".bible", "config.json")
}

func loadConfig() *Config {
	data, err := os.ReadFile(configPath())
	if err != nil {
		return &Config{Default: []string{"kjv"}}
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil || len(cfg.Default) == 0 {
		return &Config{Default: []string{"kjv"}}
	}
	return &cfg
}

func saveConfig(cfg *Config) error {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".bible")
	os.MkdirAll(dir, 0755)
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(configPath(), data, 0644)
}

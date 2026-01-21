package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const CONFIG_FILE_NAME = ".bloggatorconfig.json"

func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Cannot locate home directory: %v", err)
	}
	path := filepath.Join(home_dir, CONFIG_FILE_NAME)
	return path, nil
}

func write(cfg *Config) error {
	json_str, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	path, err := getConfigFilePath()
	err = os.WriteFile(path, json_str, 0666)
	return err
}

func Read() (Config, error) {
	cfg := Config{}
	path, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}
	file_content, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(file_content, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func (cfg *Config) SetUser(user_name string) error {
	cfg.CurrentUserName = user_name
	return write(cfg)
}

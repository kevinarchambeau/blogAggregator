package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := getPath()
	if err != nil {
		return Config{}, fmt.Errorf("unable to get config path")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("file not loaded")
	}
	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Printf("Error decoding config file: %s", err)
		return Config{}, err
	}

	return config, nil
}

func (cfg *Config) SetUser(name string) error {
	path, err := getPath()
	if err != nil {
		return fmt.Errorf("unable to get config path")
	}

	cfg.CurrentUserName = name

	data, err := json.Marshal(cfg)
	if err != nil {
		log.Printf("Error marshalling config: %s", err)
		return err
	}

	err = os.WriteFile(path, data, 0666)
	if err != nil {
		log.Printf("Error writing file: %s", err)
		return err
	}
	return nil
}

func getPath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home not found")
	}
	return path + "/.gatorconfig.json", nil
}

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
	path, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("home not found")
	}

	data, err := os.ReadFile(path + "/.gatorconfig.json")
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

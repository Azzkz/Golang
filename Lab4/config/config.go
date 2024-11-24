package config

import (
	"encoding/json"
	"os"
)

// Config struct with fields to be loaded from the config file
type Config struct {
	ServerPort  string `json:"server_port"`
	LogLevel    string `json:"log_level"`
	TLSCertPath string `json:"tls_cert_path"`
	TLSKeyPath  string `json:"tls_key_path"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

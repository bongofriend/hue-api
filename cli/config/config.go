package config

import (
	_ "embed"
	"encoding/json"
)

//go:embed config.json
var configFile []byte

type ClientConfig struct {
	ApiUrl   string `json:"apiUrl"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetConfig() (ClientConfig, error) {
	var cfg ClientConfig
	if err := json.Unmarshal(configFile, &cfg); err != nil {
		return ClientConfig{}, err
	}
	return cfg, nil
}

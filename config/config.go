package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	Port int        `json:"port"`
	Auth AuthConfig `json:"auth"`
	Hue  HueConfig  `json:"hue"`
}

type AuthConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type HueConfig struct {
	BridgeName    string `json:"bridgename"`
	PlainUsername string `json:"plainUsername"`
	AppName       string `json:"appname"`
	Username      string `json:"username"`
}

func GetConfig(path string) (AppConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return AppConfig{}, err
	}

	decoder := json.NewDecoder(file)
	var config AppConfig
	if err := decoder.Decode(&config); err != nil {
		return AppConfig{}, err
	}

	return config, nil
}

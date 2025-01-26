package services

import (
	"encoding/json"
	"fmt"
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
	BridgeName    string  `json:"bridgename"`
	PlainUsername string  `json:"plainUsername"`
	AppName       string  `json:"appname"`
	Username      *string `json:"username"`
}

type ConfigService interface {
	GetConfig() (AppConfig, error)
	UpdateConfig(config AppConfig) error
}

type configService struct {
	configFilePath string
	appConfig      AppConfig
	configRead     bool
}

func NewConfigServuce(configFilePath string) ConfigService {
	return &configService{
		configFilePath: configFilePath,
		configRead:     false,
	}
}

// GetConfig implements ConfigService.
func (c *configService) GetConfig() (AppConfig, error) {
	if c.configRead {
		return c.appConfig, nil
	}
	cfg, err := GetConfig(c.configFilePath)
	if err != nil {
		return AppConfig{}, fmt.Errorf("reading config: %w", err)
	}
	c.appConfig = cfg
	return c.appConfig, nil
}

// UpdateConfig implements ConfigService.
func (c *configService) UpdateConfig(config AppConfig) error {
	file, err := os.Open(c.configFilePath)
	if err != nil {
		return fmt.Errorf("updating config: %w", err)
	}
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("updating config: %w", err)
	}
	newConfig, err := GetConfig(c.configFilePath)
	if err != nil {
		return fmt.Errorf("updating config: %w", err)
	}
	c.appConfig = newConfig
	return nil
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

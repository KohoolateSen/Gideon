package config

import (
	"Gideon/internal/utils"
)

// Global config parameter that can be used in other packages
var AppConfig Config

// YAML fields
type ProxyConfig struct {
	Enabled    bool   `yaml:"enabled"`
	Url        string `yaml:"url"`
	NoProxyUrl string `yaml:"no_proxy"`
}

type LoggingConfig struct {
	Enabled bool   `yaml:"enabled"`
	Level   string `yaml:"level"`
}

type NewsReaderConfig struct {
	APIKey string `yaml:"api_key"`
}

type TelegramBotConfig struct {
	APIToken    string `yaml:"api_token"`
	APIEndPoint string `yaml:"api_endpoint"`
}

type Config struct {
	Proxy       ProxyConfig       `yaml:"proxy"`
	Logging     LoggingConfig     `yaml:"logging"`
	NewsReader  NewsReaderConfig  `yaml:"news_reader"`
	TelegramBot TelegramBotConfig `yaml:"telegram_bot"`
}

// Config loader
func (c *Config) LoadConfig(paths ...string) error {
	path := "./internal/config/config.yaml"

	if len(paths) > 0 {
		path = paths[0]
	}

	if err := utils.YAMLParse(path, &c); err != nil {
		return err
	}

	return nil
}

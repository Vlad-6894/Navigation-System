package log_core_logger_config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type AuthLoggerConfig struct {
	Level  string `envconfig:"LEVEL" default:"DEBUG"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func NewAuthLoggerConfig() (AuthLoggerConfig, error) {
	var config AuthLoggerConfig

	if err := envconfig.Process("LOGGER_AUTH", &config); err != nil {
		return AuthLoggerConfig{}, fmt.Errorf("fail process %w: ", err)
	}

	return config, nil
}

func NewAuthLoggerConfigMust() AuthLoggerConfig {
	config, err := NewAuthLoggerConfig()
	if err != nil {
		panic(err)
	}

	return config
}

func (l AuthLoggerConfig) GetLevel() string {
	return l.Level
}

func (l AuthLoggerConfig) GetFolder() string {
	return l.Folder
}

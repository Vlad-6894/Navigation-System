package log_core_logger_config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type GeoLoggerConfig struct {
	Level  string `envconfig:"LEVEL" default:"DEBUG"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func NewGeoLoggerConfig() (GeoLoggerConfig, error) {
	var config GeoLoggerConfig

	if err := envconfig.Process("LOGGER_GEO", &config); err != nil {
		return GeoLoggerConfig{}, fmt.Errorf("fail process %w: ", err)
	}

	return config, nil
}

func NewGeoLoggerConfigMust() GeoLoggerConfig {
	config, err := NewGeoLoggerConfig()
	if err != nil {
		panic(err)
	}

	return config
}

func (l GeoLoggerConfig) GetLevel() string {
	return l.Level
}

func (l GeoLoggerConfig) GetFolder() string {
	return l.Folder
}

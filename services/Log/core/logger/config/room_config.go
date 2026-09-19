package log_core_logger_config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type RoomLoggerConfig struct {
	Level  string `envconfig:"LEVEL" default:"DEBUG"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func NewRoomLoggerConfig() (RoomLoggerConfig, error) {
	var config RoomLoggerConfig

	if err := envconfig.Process("LOGGER_ROOM", &config); err != nil {
		return RoomLoggerConfig{}, fmt.Errorf("fail process %w: ", err)
	}

	return config, nil
}

func NewRoomLoggerConfigMust() RoomLoggerConfig {
	config, err := NewRoomLoggerConfig()
	if err != nil {
		panic(err)
	}

	return config
}

func (l RoomLoggerConfig) GetLevel() string {
	return l.Level
}

func (l RoomLoggerConfig) GetFolder() string {
	return l.Folder
}

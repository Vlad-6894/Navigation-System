package log_core_redis

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type LogRedisConfig struct {
	RedisHost     string        `envconfig:"HOST" required:"true"`
	RedisPort     string        `envconfig:"PORT" required:"true"`
	RedisPassword string        `envconfig:"PASSWORD" required:"true"`
	RedisTimeout  time.Duration `envconfig:"TIMEOUT" required:"true"`
}

func NewRedisConfig() (LogRedisConfig, error) {
	var config LogRedisConfig

	if err := envconfig.Process("LOG_REDIS", &config); err != nil {
		return LogRedisConfig{}, fmt.Errorf("fail to process redis config: %w", err)
	}

	return config, nil
}

func NewRedisConfigMust() LogRedisConfig {
	config, err := NewRedisConfig()
	if err != nil {
		panic(err)
	}

	return config
}

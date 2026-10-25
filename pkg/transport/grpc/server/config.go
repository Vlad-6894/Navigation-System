package pkg_grpc_server

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type GRPCServerConfig struct {
	Addr            string        `envconfig:"PORT" required:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
}

func NewGRPCServerConfig() (GRPCServerConfig, error) {
	var config GRPCServerConfig

	if err := envconfig.Process("GRPC_SERVER", &config); err != nil {
		return GRPCServerConfig{}, err
	}

	return config, nil
}

func NewHTTPServerConfigMust() GRPCServerConfig {
	config, err := NewGRPCServerConfig()
	if err != nil {
		panic(err)
	}

	return config
}

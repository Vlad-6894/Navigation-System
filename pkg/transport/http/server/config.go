package pkg_http_server

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type HTTPServerConfig struct {
	Addr            string        `envconfig:"PORT" required:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
}

func NewHTTPServerConfig() (HTTPServerConfig, error) {
	var config HTTPServerConfig

	if err := envconfig.Process("HTTP_SERVER", &config); err != nil {
		return HTTPServerConfig{}, err
	}

	return config, nil
}

func NewHTTPServerConfigMust() HTTPServerConfig {
	config, err := NewHTTPServerConfig()
	if err != nil {
		panic(err)
	}

	return config
}

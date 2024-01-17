package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Api struct {
	Name              string        `default:"go-app"`
	Host              string        `default:"0.0.0.0"`
	Port              string        `default:"3080"`
	RPCPort           int           `default:"50051"`
	ReadHeaderTimeout time.Duration `split_words:"true" default:"60s"`
	GracefulTimeout   time.Duration `split_words:"true" default:"8s"`

	RequestLog bool   `split_words:"true" default:"false"`
	RunSwagger bool   `split_words:"true" default:"true"`
	AuthType   string `default:"jwt"`
	JWTSecret  string `default:"hey-yok"`
}

func API() Api {
	var api Api
	envconfig.MustProcess("API", &api)

	return api
}

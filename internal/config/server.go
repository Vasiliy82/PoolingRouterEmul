package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const AppName string = "PoolingRouterEmul"

type Server struct {
	Domain   string `envconfig:"DOMAIN"`
	TraceURL string `envconfig:"TRACE_URL"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"WARN"`
	appName  string

	ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"5s"`
	WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"5s"`
	Port         string        `envconfig:"PORT" default:":80"`
}

func loadServer(appPrefix string) (*Server, error) {
	var a Server

	err := envconfig.Process(appPrefix, &a)
	if err != nil {
		return nil, err
	}

	if a.Domain == "" {
		a.appName = AppName

	} else {
		a.appName = a.Domain + ":" + AppName
	}

	return &a, nil
}

func (a *Server) AppName() string {
	return a.appName
}

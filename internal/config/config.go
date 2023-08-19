package config

import (
	"github.com/pkg/errors"
)

const appNamePrefix = "POOLING_RTR"

var approveDebugNoAuth bool // = true

type Config struct {
	Server *Server
}

func LoadAll() (*Config, error) {
	confApp, err := loadServer(appNamePrefix)
	if err != nil {
		return nil, errors.Wrap(err, "error loadApp")
	}

	return &Config{
		Server: confApp,
	}, nil
}

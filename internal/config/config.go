package config

import (
	"github.com/pkg/errors"
)

const appNamePrefix = "POOLING_RTR"

var approveDebugNoAuth bool // = true

type Config struct {
	App *App
}

func LoadAll() (*Config, error) {
	confApp, err := loadApp(appNamePrefix)
	if err != nil {
		return nil, errors.Wrap(err, "error loadApp")
	}

	return &Config{
		App: confApp,
	}, nil
}

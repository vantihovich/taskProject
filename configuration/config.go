package config

import (
	"github.com/caarlos0/env/v6"
)

type Database struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE"`
}

type App struct {
	Database Database
}

func Load() (App, error) {

	cfg := App{}
	if err := env.Parse(&cfg); err != nil {

		return App{}, err
	}
	return cfg, nil
}

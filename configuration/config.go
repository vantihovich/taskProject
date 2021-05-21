package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v6"
)

type Database struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE"`
}

type App struct {
	Database Database
}

//func Load(cfg App, err error) (string, error) {

func Load() (cfgDb string, err error) {

	cfg := App{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	cff := fmt.Sprintf("%+v", cfg.Database)

	cff2 := strings.TrimPrefix(cff, "{")
	cfgDb = strings.TrimSuffix(cff2, "}")

	fmt.Println("The configs are:", cfgDb)
	return cfgDb, err
}

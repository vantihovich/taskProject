//package config
package main

import (
	"fmt"

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

//func Load(App, err) {
func main() {

	cfg := App{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	cffg := fmt.Sprintf("%+v", cfg.Database)
	fmt.Println("Struct with configs", cfg)

	fmt.Println("The configs are:", cffg)
}

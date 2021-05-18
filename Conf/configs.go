package conf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Configs() (cfg Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println("Error opening the config file", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Error decoding the config file", err)
	}
	fmt.Println("configs are:", cfg)
	return cfg
}

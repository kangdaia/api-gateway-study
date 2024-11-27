package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App []App `yaml:"apps"`
}

type App struct {
	App struct {
		Port    string `yaml:"port"`
		Version string `yaml:"version"`
		Name    string `yaml:"name"`
	} `yaml:"app"`

	Http     HttpCfg  `yaml:"http"`
	Producer Producer `yaml:"kafka"`
}

func NewCfg(path string) Config {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var c Config
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		panic(err)
	}
	return c
}

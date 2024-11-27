package config

import "api-gateway-study/types/http"

type HttpCfg struct {
	Router  []Router `yaml:"router"`
	BaseUrl string   `yaml:"base_url"`
}

type Router struct {
	Method   http.HttpMethod `yaml:"method"`
	GetType  http.GetType    `yaml:"get_type"`
	Variable []string        `yaml:"variable"`
	Path     string          `yaml:"path"`

	Auth   *Auth             `yaml:"auth"`
	Header map[string]string `yaml:"header"`
}

type Auth struct {
	Key   string `yaml:"key"`
	Token string `yaml:"token"`
}

package config

type AppConfig struct {
	Env  string `yaml:"env"`
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

var App AppConfig

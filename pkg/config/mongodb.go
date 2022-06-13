package config

type MongoConfig struct {
	TimeoutSecond int    `yaml:"timeout_second"`
	ConnectUrl    string `yaml:"connect_url"`
	DatabaseName  string `yaml:"database_name"`
}

var Mongo MongoConfig

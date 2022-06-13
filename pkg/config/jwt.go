package config

type JwtConfig struct {
	Secret          string `yaml:"secret"`
	ExpiredSeconds  int64  `yaml:"expired_seconds"`
	Issuer          string `yaml:"issuer"`
	GraphMeEndpoint string `yaml:"graph_me_endpoint"`
	AdminKey        string `yaml:"admin_key"`
}

var Jwt JwtConfig

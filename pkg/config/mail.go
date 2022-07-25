package config

type MailConfig struct {
	Smtp struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"smtp"`

	From struct {
		Address string `yaml:"address"`
		Name    string `yaml:"name"`
	} `yaml:"from"`
}

var Mail MailConfig

package bootstrap

import (
	"io/ioutil"
	"log"
	"path"
	"projmural-backend/pkg/config"

	"gopkg.in/yaml.v2"
)

func LoadEnv(env string) {
	basePath := path.Join(config.EnvPos, env)

	appPath := path.Join(basePath, "app.yml")
	loadConfig(appPath, &config.App)
	if config.App.Env != env {
		log.Panicln("environment variables conflict")
	}
	log.Println(config.App)

	mongodbPath := path.Join(basePath, "mongodb.yml")
	loadConfig(mongodbPath, &config.Mongo)
	log.Println(config.Mongo)

	jwtPath := path.Join(basePath, "jwt.yml")
	loadConfig(jwtPath, &config.Jwt)
	log.Println(config.Jwt)
}

func loadConfig(path string, config interface{}) {
	envfile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(envfile, config)
	if err != nil {
		log.Println(err)
	}
}

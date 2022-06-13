package main

import (
	"flag"
	"log"
	"projmural-backend/bootstrap"
	"projmural-backend/dao"
	"projmural-backend/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//--env=local
	var env string
	flag.StringVar(&env, "env", "local", "load environment variables, --env=local meaning env=local")
	flag.Parse()
	log.Println(env)
	// load env
	bootstrap.LoadEnv(env)

	dao.NewMongoDao() //init mongodb
	r := gin.Default()
	api := r.Group("/api")
	http.Init(api) //register route
	r.Run(":8081")
}

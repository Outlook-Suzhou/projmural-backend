package main

import (
	"projmural-backend/dao"
	"projmural-backend/http"

	"github.com/gin-gonic/gin"
)

func main() {
	dao.NewMongoDao() //init mongodb
	r := gin.Default()
	api := r.Group("/api")
	http.Init(api) //register route
	r.Run(":8081")
}

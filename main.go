package main

import (
	"projmural-backend/dao"
	"projmural-backend/http"

	"github.com/gin-gonic/gin"
)

func main() {
	dao.NewMongoDao()
	r := gin.Default()
	api := r.Group("/api")
	http.Init(api)
	r.Run(":8081")
}
package main

import (
	"projmural-backend/dao"
	"projmural-backend/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dao.NewMongoDao()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	http.Init(r)
	r.Run(":8081")
}
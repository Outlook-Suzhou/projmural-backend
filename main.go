package main

import (
	"github.com/gin-gonic/gin"
	"projmural-backend/dao"
	"projmural-backend/http"
)

func main() {
	dao.Init()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	http.Init(r)
	r.Run(":8081")
}
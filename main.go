package main

import (
	"github.com/gin-gonic/gin"
	"projmural-backend/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	http.Init(r)
	r.Run(":8081")
}
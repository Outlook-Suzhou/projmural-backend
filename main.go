package main

import (
<<<<<<< HEAD
=======
	"github.com/gin-gonic/gin"
>>>>>>> main
	"projmural-backend/dao"
	"projmural-backend/http"

	"github.com/gin-gonic/gin"
)

func main() {
	dao.Init()
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
package http

import "github.com/gin-gonic/gin"

func user (ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
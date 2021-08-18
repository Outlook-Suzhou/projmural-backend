package http

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	r.POST("/login", login)
}
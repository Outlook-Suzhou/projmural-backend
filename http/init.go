package http

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	r.POST("/login", jsonParserMiddleWare(),  login)
	r.POST("/user", jsonParserMiddleWare(), user)
}
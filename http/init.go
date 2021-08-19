package http

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	r.POST("/login", RouterMiddleWare(login))
	r.POST("/user", JwtMiddleWare(), user);
}
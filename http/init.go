package http

import "github.com/gin-gonic/gin"

func Init(r *gin.RouterGroup) {
	r.POST("/login", jsonParserMiddleWare(), login)
	r.POST("/user", jwtMiddleWare(), jsonParserMiddleWare(), user)
	r.POST("/sendmail", jwtMiddleWare(), jsonParserMiddleWare(), SendEmail)
	r.GET("/currentUser", jwtMiddleWare(), currentUser)
}

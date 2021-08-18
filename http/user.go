package http

import "github.com/gin-gonic/gin"

func user (getBody GetBodyFunction, claims *Claims) (int, *gin.H){

	return RESP_OK, &gin.H{}
}
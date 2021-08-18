package http

import "github.com/gin-gonic/gin"

const (
	RESP_OK = 1
	RESP_SERVER_ERROR = 2
)

func quickResp(cmd int, ctx *gin.Context){
	switch cmd {
	case RESP_OK:
		ctx.JSON(200, gin.H{
			"msg": "ok",
			"retc": 0,
		})
	case RESP_SERVER_ERROR:
		ctx.JSON(500, gin.H{
			"msg": "server error",
			"retc": -2,
		})
	}
}

func okRespWithData(ctx *gin.Context, data *gin.H){
	ctx.JSON(200, gin.H{
		"msg": "ok",
		"retc": 0,
		"data": data,
	})
}
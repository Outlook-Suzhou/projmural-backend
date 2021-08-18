package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

const (
	RESP_OK_WITH_DATA = 0
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

type routerFunction func(ctx *gin.Context)
type coreFunction func(ctx *gin.Context) (int, error, *gin.H)

func routerMiddleWare(core coreFunction) routerFunction {
	return func(ctx *gin.Context) {
		retc, err, data := core(ctx)
		if retc == RESP_OK_WITH_DATA{
			okRespWithData(ctx, data)
		} else {
			quickResp(retc, ctx)
		}
		if err != nil {
			log.Println(err)
		}
	}
}
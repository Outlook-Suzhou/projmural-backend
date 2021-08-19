package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

const (
	RESP_OK_WITH_DATA = 1
	RESP_OK = 0
	RESP_JWT_FAIL = -1
	RESP_SERVER_ERROR = -2
	RESP_ACCESS_TOKEN_FAIL = -3
)

func quickResp(cmd int, ctx *gin.Context){
	switch cmd {
	case RESP_OK:
		ctx.JSON(200, gin.H{
			"msg": "ok",
			"retc": cmd,
		})
	case RESP_JWT_FAIL:
		ctx.JSON(200, gin.H{
			"msg": "jwt fail",
			"retc": cmd,
		})
	case RESP_SERVER_ERROR:
		ctx.JSON(500, gin.H{
			"msg": "server error",
			"retc": cmd,
		})
	case RESP_ACCESS_TOKEN_FAIL:
		ctx.JSON(200, gin.H{
			"msg": "access token fail",
			"retc": cmd,
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

type CoreFunction func(GetBodyFunction) (int, *gin.H)
type CoreJwtFunction func(function GetBodyFunction, claims *Claims) (int, *gin.H)
type GetBodyFunction func(interface{})

func JwtMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt, exsit := ctx.Request.Header["Authorization"]
		if exsit == true {
			var c *Claims
			c, err := ParseJWT(jwt[0][7:])
			if err == nil {
				// retc, data := core(getBody, c)
				ctx.Next();
			} else {
				quickResp(RESP_JWT_FAIL, ctx)
				panic(err)
			}
		} else {
			quickResp(RESP_JWT_FAIL, ctx)
		}
		return
	}
}

func jsonRequestMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var getBody = func(i interface{}) {
			data, err := ioutil.ReadAll(ctx.Request.Body)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(data, i)
			if err != nil {
				panic(err)
			}
		}
		ctx.Set("getBody", getBody)
		ctx.Next()
	}
}
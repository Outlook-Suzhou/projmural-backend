package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

const (
	RESP_OK_WITH_DATA = 1
	RESP_OK = 0
	RESP_JWT_FAIL = -1
	RESP_SERVER_ERROR = -2
	RESP_ACCESS_TOKEN_FAIL = -3
	RESP_USER_NOT_EXIST = -4
	RESP_INVALID_OPERATION = -5
	RESP_INVALID_JSON_FORMAT = -6
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
	case RESP_USER_NOT_EXIST:
		ctx.JSON(200, gin.H{
			"msg": "user not exist",
			"retc": cmd,
		})
	case RESP_INVALID_OPERATION:
		ctx.JSON(200, gin.H{
			"msg": "invalid operation",
			"retc": cmd,
		})
	case RESP_INVALID_JSON_FORMAT:
		ctx.JSON(200, gin.H{
			"msg": "invalid json format",
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

func jwtMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt, exsit := ctx.Request.Header["Authorization"]
		if exsit == true {
			var c *Claims
			c, err := ParseJWT(jwt[0][7:])
			ctx.Set("claim", c)
			fmt.Println(c)
			if err == nil {
				ctx.Next()
			} else {
				quickResp(RESP_JWT_FAIL, ctx)
				ctx.Abort()
				return
			}
		} else {
			quickResp(RESP_JWT_FAIL, ctx)
			ctx.Abort()
		}
		return
	}
}

func jsonParserMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var getBody GetBodyFunction
		getBody = func(i interface{}) {
			data, err := ioutil.ReadAll(ctx.Request.Body)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(data, i)
			if err != nil {
				quickResp(RESP_INVALID_JSON_FORMAT, ctx)
				ctx.Abort()
				panic(err)
			}
		}
		ctx.Set("getBody", getBody)
		ctx.Next()
	}
}
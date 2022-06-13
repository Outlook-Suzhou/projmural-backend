package http

import (
	"encoding/json"
	"io/ioutil"
	"projmural-backend/pkg/config"

	"github.com/gin-gonic/gin"
)

const (
	RESP_OK_WITH_DATA        = 1
	RESP_OK                  = 0
	RESP_JWT_FAIL            = -1
	RESP_SERVER_ERROR        = -2
	RESP_ACCESS_TOKEN_FAIL   = -3
	RESP_USER_NOT_EXIST      = -4
	RESP_INVALID_OPERATION   = -5
	RESP_INVALID_JSON_FORMAT = -6
	RESP_PERMISSION_DENY     = -7
)

var respMsg = map[int]string{
	RESP_OK:                  "ok",
	RESP_JWT_FAIL:            "jwt fail",
	RESP_SERVER_ERROR:        "server error",
	RESP_ACCESS_TOKEN_FAIL:   "access token fail",
	RESP_USER_NOT_EXIST:      "user not exist",
	RESP_INVALID_OPERATION:   "invalid operation",
	RESP_INVALID_JSON_FORMAT: "invalid json format",
	RESP_PERMISSION_DENY:     "permission deny",
}

func quickResp(cmd int, ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg":  respMsg[cmd],
		"retc": cmd,
	})
}

func okRespWithData(ctx *gin.Context, data *gin.H) {
	ctx.JSON(200, gin.H{
		"msg":  "ok",
		"retc": 0,
		"data": data,
	})
}

type CoreFunction func(GetBodyFunction) (int, *gin.H)
type CoreJwtFunction func(function GetBodyFunction, claims *Claims) (int, *gin.H)
type GetBodyFunction func(interface{})

//get jwt from header
func jwtMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt, exsit := ctx.Request.Header["Authorization"]
		if exsit == true {
			if jwt[0] == config.Jwt.AdminKey {
				c := &Claims{MicrosoftId: "admin"}
				ctx.Set("claim", c)
				ctx.Next()
				return
			}
			var c *Claims
			c, err := ParseJWT(jwt[0][7:]) //removed bearer
			ctx.Set("claim", c)
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
			// if err != nil {
			// panic(err)
			// }
			//fmt.Printf("data:[%v]\nerr:[%v]\n", string(data), err)
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

package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"projmural-backend/dao"
)

type loginRequest struct {
	Data dao.User	`json:"data"`
}

func login(ctx *gin.Context) {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		quickResp(RESP_SERVER_ERROR, ctx)
		return
	}
	log.Println(string(data))
	var request loginRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		log.Println(err)
		quickResp(RESP_SERVER_ERROR, ctx)
		return
	}
	jwt, err := GenerateJWT(request.Data.MicrosoftId)
	if err != nil {
		log.Println(err)
		quickResp(RESP_SERVER_ERROR, ctx)
		return
	}
	okRespWithData(ctx, &gin.H{
		"jwt": jwt,
	})
}
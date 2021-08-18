package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type loginRequest struct {
	AccessToken string `json:"access_token"`
}

func login(ctx *gin.Context) (int, error, *gin.H) {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return RESP_SERVER_ERROR, err, &gin.H{}
	}
	log.Println(string(data))
	var request loginRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		return RESP_SERVER_ERROR, err, &gin.H{}
	}
	return RESP_OK, nil, &gin.H{}
}
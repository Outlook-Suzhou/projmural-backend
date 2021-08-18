package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type loginRequest struct {
	AccessToken string `json:"access_token"`
}

type graphRespond struct {
	MicrosoftId string `json:"id"`
}

func login(ctx *gin.Context) (int, error, *gin.H) {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return RESP_SERVER_ERROR, err, &gin.H{}
	}
	var request loginRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		return RESP_SERVER_ERROR, err, &gin.H{}
	}
	client := &http.Client{}
	getRequest, _ := http.NewRequest("GET", GRAPH_ME_ENDPOINT, nil)
	getRequest.Header.Add("Authorization", "Bearer " + request.AccessToken)
	respond, err := client.Do(getRequest)
	if err != nil {
		return RESP_SERVER_ERROR, err, &gin.H{}
	}
	defer respond.Body.Close()
	body, err := ioutil.ReadAll(respond.Body)
	if err != nil {
		return RESP_SERVER_ERROR, err, &gin.H{}
	}
	var graphResp graphRespond
	json.Unmarshal(body, &graphResp)
	log.Println(graphResp)
	if graphResp.MicrosoftId == "" {return RESP_ACCESS_TOKEN_FAIL, nil, &gin.H{}}
	jwt, err := GenerateJWT(graphResp.MicrosoftId)
	return RESP_OK_WITH_DATA, nil, &gin.H{"jwt": jwt}
}
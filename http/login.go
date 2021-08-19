package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type LoginRequest struct {
	AccessToken string `json:"access_token"`
}

type GraphRespond struct {
	MicrosoftId string `json:"id"`
}

func login(ctx *gin.Context) {
	getBodyInterface, has := ctx.Get("getBody")
	getBody := getBodyInterface.(GetBodyFunction)
	if has == false {
		quickResp(RESP_SERVER_ERROR, ctx)
		panic("getBody is not exist")
		return
	}
	var request LoginRequest
	getBody(&request)
	client := &http.Client{}
	getRequest, _ := http.NewRequest("GET", GRAPH_ME_ENDPOINT, nil)
	getRequest.Header.Add("Authorization", "Bearer " + request.AccessToken)
	respond, err := client.Do(getRequest)
	if err != nil {panic(err)}
	defer respond.Body.Close()
	body, err := ioutil.ReadAll(respond.Body)
	if err != nil {panic(err)}
	var graphResp GraphRespond
	err = json.Unmarshal(body, &graphResp)
	if err != nil {panic(err)}
	if graphResp.MicrosoftId == "" {quickResp(RESP_ACCESS_TOKEN_FAIL, ctx); return}
	jwt, err := GenerateJWT(graphResp.MicrosoftId)
	okRespWithData(ctx, &gin.H{"jwt": jwt})
	return
}
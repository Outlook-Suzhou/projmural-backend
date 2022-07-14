package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"projmural-backend/dao"
	"projmural-backend/pkg/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginRequest struct {
	AccessToken string `json:"access_token"`
}

type GraphRespond struct {
	MicrosoftId string `json:"id"`
	Name        string `json:"displayName"`
	Mail        string `json:"mail"`
}

func login(ctx *gin.Context) {
	// check method exist
	getBodyInterface, has := ctx.Get("getBody")
	getBody := getBodyInterface.(GetBodyFunction)
	if has == false {
		quickResp(RESP_SERVER_ERROR, ctx)
		panic("getBody is not exist")
	}

	var request LoginRequest
	getBody(&request)

	//send request to GRAPH_ME_ENDPOINT to get MicrosoftId
	client := &http.Client{}
	getRequest, _ := http.NewRequest("GET", config.Jwt.GraphMeEndpoint, nil)
	getRequest.Header.Add("Authorization", "Bearer "+request.AccessToken)
	respond, err := client.Do(getRequest)
	if err != nil {
		panic(err)
	}
	defer respond.Body.Close()
	body, err := ioutil.ReadAll(respond.Body)
	if err != nil {
		panic(err)
	}
	var graphResp GraphRespond
	err = json.Unmarshal(body, &graphResp)
	if err != nil {
		panic(err)
	}
	if graphResp.MicrosoftId == "" {
		quickResp(RESP_ACCESS_TOKEN_FAIL, ctx)
		return
	}

	// check database
	dataBase := dao.GetMongoDao()
	_, err = dataBase.FindUserByMicrosoftId(graphResp.MicrosoftId)
	if err == mongo.ErrNoDocuments {
		dataBase.InsertOrReplaceUserByMicrosoftId(dao.User{
			MicrosoftId: graphResp.MicrosoftId,
			Name:        graphResp.Name,
			Mail:        "",
			Photo:		 "",
			Canvas:      []dao.CanvaInfo{},
			RecentCanvas:      []dao.RecentCanvaInfo{},
			Tasks:      []dao.TaskInfo{},
		})
	}

	//GenerateJWT
	jwt, err := GenerateJWT(graphResp.MicrosoftId)
	okRespWithData(ctx, &gin.H{"jwt": jwt})
}

package http

import (
	"fmt"
	"projmural-backend/dao"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Type string   `json:"type"`
	Data dao.User `json:"data"`
}

func user(ctx *gin.Context) {
	getBodyInterface, has := ctx.Get("getBody")
	getBody := getBodyInterface.(GetBodyFunction)
	if has == false {
		quickResp(RESP_SERVER_ERROR, ctx)
		panic("getBody is not exist")
		return
	}
	var request UserRequest
	var dataBase *dao.MongoDao = dao.GetMongoDao()
	getBody(&request)
	switch request.Type {
	case "query":
		dataBase.FindUserByMicrosoftId(request.Data.MicrosoftId)
	case "update":
		fmt.Println(request)
		dataBase.InsertOrReplaceUserByMicrosoftId(request.Data)
	case "insert":
		dataBase.InsertOrReplaceUserByMicrosoftId(request.Data)
	}


}

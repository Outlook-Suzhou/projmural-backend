package http

import (
	"fmt"
	"projmural-backend/dao"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
		user, err := dataBase.FindUserByMicrosoftId(request.Data.MicrosoftId)
		if(err == mongo.ErrNoDocuments) {
			quickResp(RESP_USER_NOT_EXIST, ctx)
			return;
		} else {
			okRespWithData(ctx, user.(interface{}))
		}
	case "update":
		fmt.Println(request)
		dataBase.InsertOrReplaceUserByMicrosoftId(request.Data)
	case "insert":
		dataBase.InsertOrReplaceUserByMicrosoftId(request.Data)
	}


}

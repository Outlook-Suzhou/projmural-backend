package http

import (
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
		if err == mongo.ErrNoDocuments {
			quickResp(RESP_USER_NOT_EXIST, ctx)
			return
		} else if err == nil {
			okRespWithData(ctx, user.GinH())
			return
		} else {
			panic(err)
		}
	case "update":
		dataBase.InsertOrReplaceUserByMicrosoftId(request.Data)
	case "insert":
		dataBase.InsertOrReplaceUserByMicrosoftId(request.Data)
	default:
		quickResp(RESP_INVALID_OPERATION, ctx)
		return
	}
	quickResp(RESP_OK, ctx)
}

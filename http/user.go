package http

import (
	"errors"
	"projmural-backend/dao"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRequest struct {
	Type string   `json:"type"` // operator type update/insert/query
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
	getBody(&request)
	claimInterface, has := ctx.Get("claim")
	if has == false {
		errors.New("claim is not exist")
	}
	claim := claimInterface.(*Claims)

	//check MicrosoftId in body and jwt
	if claim.MicrosoftId != "admin" && claim.MicrosoftId != request.Data.MicrosoftId {
		quickResp(RESP_PERMISSION_DENY, ctx)
		return
	}
	var dataBase *dao.MongoDao = dao.GetMongoDao()

	//do db operation based on request.Type
	switch request.Type {
	case "queryall":
		users, err := dataBase.FindAllUsers()
		if err == mongo.ErrNoDocuments {
			quickResp(RESP_USER_NOT_EXIST, ctx)
			return
		} else if err == nil {
			ginh_users := gin.H{"users": users}
			okRespWithData(ctx, &ginh_users)
			return
		} else {
			panic(err)
		}
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

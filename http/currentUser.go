package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"projmural-backend/dao"
)

func currentUser(ctx *gin.Context) {
	claimsInterface, has := ctx.Get("claim")
	if has == false {panic(errors.New("claim is not exist"))}
	claim := claimsInterface.(*Claims)
	var dataBase *dao.MongoDao = dao.GetMongoDao()
	user, err := dataBase.FindUserByMicrosoftId(claim.MicrosoftId)
	if err == mongo.ErrNoDocuments {
		quickResp(RESP_USER_NOT_EXIST, ctx)
		return
	} else if err != nil {panic(err)}
	okRespWithData(ctx, user.GinH())
	return
}
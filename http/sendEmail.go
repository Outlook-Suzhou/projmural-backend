package http

import (
	"errors"
	"projmural-backend/dao"
	"projmural-backend/pkg/config"
	"projmural-backend/pkg/mail"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type MailRequest struct {
	Id      string `json:"id"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func send(email string, subject string, content string) error {
	// 发送邮件
	suc := mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: config.Mail.From.Address,
			Name:    config.Mail.From.Name,
		},
		To:      []string{email},
		Subject: subject,
		HTML:    []byte(content),
	})
	if !suc {
		return errors.New("send email failed")
	}
	return nil
}

func SendEmail(ctx *gin.Context) {
	getBodyInterface, has := ctx.Get("getBody")
	getBody := getBodyInterface.(GetBodyFunction)
	if has == false {
		quickResp(RESP_SERVER_ERROR, ctx)
		panic("getBody is not exist")
		return
	}
	var request MailRequest
	getBody(&request)
	claimInterface, has := ctx.Get("claim")
	if has == false {
		errors.New("claim is not exist")
	}
	claim := claimInterface.(*Claims)
	//check MicrosoftId in body and jwt
	if claim.MicrosoftId != "admin" && claim.MicrosoftId != request.Id {
		quickResp(RESP_PERMISSION_DENY, ctx)
		return
	}

	var dataBase *dao.MongoDao = dao.GetMongoDao()
	user, err := dataBase.FindUserByMicrosoftId(request.Id)
	if err == mongo.ErrNoDocuments {
		quickResp(RESP_USER_NOT_EXIST, ctx)
	} else if err == nil {
		// okRespWithData(ctx, user.GinH())
		err := send(user.Mail, request.Subject, request.Content)
		if err != nil {
			quickResp(RESP_SERVER_ERROR, ctx)
		} else {
			quickResp(RESP_OK, ctx)
		}
	} else {
		panic(err)
	}
}

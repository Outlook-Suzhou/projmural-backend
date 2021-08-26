package dao

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CanvaInfo struct {
	ID string `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
	RecentOpen int32 `bson:"recent_open", json:"recent_open"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	MicrosoftId string             `bson:"microsoft_id,omitempty" json:"microsoft_id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Mail        string             `bson:"mail,omitempty" json:"mail"`
	Canvas      []CanvaInfo           `bson:"canvas,omitempty" json:"canvas"`
}

func (u User) Bson() bson.D {
	return bson.D{{"microsoft_id", u.MicrosoftId}, {"mail", u.Mail},
		{"name", u.Name}, {"canvas", u.Canvas}}
}

func (u User) GinH() *gin.H {
	return &gin.H{
		"microsoft_id": u.MicrosoftId,
		"name":         u.Name,
		"mail":			u.Mail,
		"canvas":       u.Canvas,
	}
}

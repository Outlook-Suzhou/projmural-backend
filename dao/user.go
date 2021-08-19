package dao

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	MicrosoftId string             `bson:"microsoft_id,omitempty" json:"microsoft_id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Canvas      []string           `bson:"canvas,omitempty" json:"canvas"`
}

func (u User) Bson() bson.D {
	return bson.D{{"microsoft_id", u.MicrosoftId},
		{"name", u.Name}, {"canvas", u.Canvas}}
}

func (u User) GinH() *gin.H {
	return &gin.H{
		"microsoft_id": u.MicrosoftId,
		"name":         u.Name,
		"canvas":       u.Canvas,
	}
}

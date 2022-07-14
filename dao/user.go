package dao

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"-"` // id in mongodb
	MicrosoftId  string             `bson:"microsoft_id,omitempty" json:"microsoft_id"`
	Name         string             `bson:"name,omitempty" json:"name"`
	Mail         string             `bson:"mail,omitempty" json:"mail"`
	Photo        string             `bson:"photo,omitempty" json:"photo"`
	Canvas       []CanvaInfo        `bson:"canvas,omitempty" json:"canvas"`
	RecentCanvas []RecentCanvaInfo  `bson:"recent_canvas,omitempty" json:"recent_canvas"`
	Tasks        []TaskInfo         `bson:"tasks,omitempty" json:"tasks"`
}

type CanvaInfo struct {
	ID   string `bson:"id" json:"id"` // more information is in sharedb, you could refer projmural-frontend
	Type string `bson:"type" json:"type"`
}

type RecentCanvaInfo struct {
	ID         string `bson:"id" json:"id"` // more information is in sharedb, you could refer projmural-frontend
	RecentOpen int64  `bson:"recent_open" json:"recent_open"`
}

type TaskInfo struct {
	ID   string `bson:"id" json:"id"` // more information is in sharedb, you could refer projmural-frontend
	Type string `bson:"type" json:"type"`
}

func (u User) Bson() bson.D {
	return bson.D{{"microsoft_id", u.MicrosoftId}, {"mail", u.Mail},
		{"name", u.Name}, {"photo", u.Photo}, {"canvas", u.Canvas}, {"recent_canvas", u.RecentCanvas}, {"tasks", u.Tasks}}
}

func (u User) GinH() *gin.H {
	return &gin.H{
		"microsoft_id":  u.MicrosoftId,
		"name":          u.Name,
		"mail":          u.Mail,
		"photo":         u.Photo,
		"canvas":        u.Canvas,
		"recent_canvas": u.RecentCanvas,
		"tasks":         u.Tasks,
	}
}

package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	MicrosoftId string             `bson:"microsoft_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Canvas      []string           `bson:"canvas,omitempty"`
}

func (u User) Bson() bson.D {
	return bson.D{{"microsoft_id", u.MicrosoftId},
		{"name", u.Name}, {"canvas", u.Canvas}}
}
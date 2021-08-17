package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	MicrosoftId string             `bson:"microsoft_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Canvas      []string           `bson:"canvas,omitempty"`
}

func main() {
	dao := NewMongoDao()
}

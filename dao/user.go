package dao

type User struct {
	ID     primitive.ObjectID `bson:"_id"`
	MicrosoftId	string `bson:"microsoft_id"`
	Name	string `bson:"name"`
	Canvas []string `bson:"canvas"`
}
package dao

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDao struct {
	mongoClient   *mongo.Client
	mongoDatabase *mongo.Database
}

func (d *MongoDao) Init() {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_SECOND*time.Second)
	defer cancel()
	var err error
	d.mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(CONNECT_URL))
	if err != nil {
		panic(err)
	}
	d.mongoDatabase = d.mongoClient.Database(DATABASE_NAME)
}

func (d *MongoDao) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_SECOND*time.Second)
	defer cancel()
	if err := d.mongoClient.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func (d MongoDao) insertUser(user User) error {
	userCollection := d.mongoDatabase.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_SECOND*time.Second)
	defer cancel()
	_, err := userCollection.InsertOne(ctx, user.Bson())
	return err
}

func (d MongoDao) updateUserByMicrosoftId(microsoftId string, user User) error {
	userCollection := d.mongoDatabase.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_SECOND*time.Second)
	defer cancel()
	filter := bson.D{{"microsoft_id", microsoftId}}
	replacement := user.Bson()
	return userCollection.FindOneAndReplace(ctx, filter, replacement).Err()
}

func (d MongoDao) InsertOrReplaceUserByMicrosoftId(user User) error {
	fmt.Println("here")
	_, err := d.FindUserByMicrosoftId(user.MicrosoftId)
	if err == mongo.ErrNoDocuments {
		return d.insertUser(user)
	}
	return d.updateUserByMicrosoftId(user.MicrosoftId, user)
}

func (d MongoDao) FindUserByMicrosoftId(microsoftId string) (User, error) {
	userCollection := d.mongoDatabase.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_SECOND*time.Second)
	defer cancel()
	res := userCollection.FindOne(ctx, bson.D{{"microsoft_id", microsoftId}})
	if res.Err() != nil {
		return User{}, res.Err()
	}
	var user User
	res.Decode(&user)
	return user, nil
}

func (d MongoDao) DeleteUserbyMicrsoftId(microsoftId string) error {
	userCollection := d.mongoDatabase.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_SECOND*time.Second)
	defer cancel()
	res := userCollection.FindOneAndDelete(ctx, bson.D{{"microsoft_id", microsoftId}})
	return res.Err()
}
var dao *MongoDao
func NewMongoDao() *MongoDao {
	dao = &MongoDao{}
	dao.Init()
	return dao
}
func GetMongoDao() *MongoDao {
	return dao;
}

package dao

import (
	"context"
	"projmural-backend/pkg/config"
	"sort"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Mongo.TimeoutSecond)*time.Second)
	defer cancel()
	var err error
	d.mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo.ConnectUrl))
	if err != nil {
		panic(err)
	}
	d.mongoDatabase = d.mongoClient.Database(config.Mongo.DatabaseName)
}

func (d *MongoDao) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Mongo.TimeoutSecond)*time.Second)
	defer cancel()
	if err := d.mongoClient.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func (d MongoDao) insertUser(user User) {
	userCollection := d.mongoDatabase.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Mongo.TimeoutSecond)*time.Second)
	defer cancel()
	_, err := userCollection.InsertOne(ctx, user.Bson())
	if err != nil {
		panic(err)
	}
}

func (d MongoDao) updateUserByMicrosoftId(microsoftId string, user User) {
	userCollection := d.mongoDatabase.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Mongo.TimeoutSecond)*time.Second)
	defer cancel()
	filter := bson.D{{"microsoft_id", microsoftId}}
	replacement := user.Bson()
	err := userCollection.FindOneAndReplace(ctx, filter, replacement).Err()
	if err != nil {
		panic(err)
	}
}

func (d MongoDao) InsertOrReplaceUserByMicrosoftId(user User) {
	_, err := d.FindUserByMicrosoftId(user.MicrosoftId)
	if err == mongo.ErrNoDocuments {
		d.insertUser(user)
	}
	d.updateUserByMicrosoftId(user.MicrosoftId, user)
}

type RecentCanvaInfoArray []RecentCanvaInfo

func (s RecentCanvaInfoArray) Len() int {
	return len(s)
}
func (s RecentCanvaInfoArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s RecentCanvaInfoArray) Less(i, j int) bool {
	return s[i].RecentOpen > s[j].RecentOpen
}

func (d MongoDao) FindAllUsers() ([]User, error) {
	userCollection := d.mongoDatabase.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Mongo.TimeoutSecond)*time.Second)
	defer cancel()
	var users []User
	cur, err := userCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return users, err
	}

	//Finding multiple documents returns a cursor
	//Iterate through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {
		//Create a value into which the single document can be decoded
		var elem User
		cur.Decode(&elem)
		users = append(users, elem)
	}
	return users, nil
}

func (d MongoDao) FindUserByMicrosoftId(microsoftId string) (User, error) {
	userCollection := d.mongoDatabase.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Mongo.TimeoutSecond)*time.Second)
	defer cancel()
	res := userCollection.FindOne(ctx, bson.D{{"microsoft_id", microsoftId}})
	if res.Err() != nil {
		return User{}, res.Err()
	}
	var user User
	res.Decode(&user)
	sort.Sort(RecentCanvaInfoArray(user.RecentCanvas))
	return user, nil
}

func (d MongoDao) DeleteUserbyMicrsoftId(microsoftId string) {
	userCollection := d.mongoDatabase.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Mongo.TimeoutSecond)*time.Second)
	defer cancel()
	res := userCollection.FindOneAndDelete(ctx, bson.D{{"microsoft_id", microsoftId}})
	if res.Err() != nil {
		panic(res.Err())
	}
}

var dao *MongoDao

func NewMongoDao() *MongoDao {
	dao = &MongoDao{}
	dao.Init()
	return dao
}
func GetMongoDao() *MongoDao {
	return dao
}

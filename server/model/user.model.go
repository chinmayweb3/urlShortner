package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/chinmayweb3/urlshortner/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UserIp     string    `json:"userIp" bson:"userIp"`
	UrlLimit   int       `json:"urlLimit" bson:"urlLimit"`
	LastViewed time.Time `json:"lastViewed" bson:"lastViewed"`
}

func FindUserByIp(uIp string) (User, error) {
	var findUser User
	if err := database.Db.Collection("users").FindOne(database.Ctx, bson.D{{Key: "userIp", Value: uIp}}).Decode(&findUser); err != nil {
		fmt.Println("no User found FindUserByIp")
		return findUser, errors.New("no User found")
	}
	return findUser, nil
}

func UserUpdate(u User) (string, error) {
	updateDoc := bson.D{{Key: "userIp", Value: u.UserIp}}
	filter := bson.D{{Key: "$set", Value: u}}
	option := options.Update().SetUpsert(true)
	up, e := database.Db.Collection("users").UpdateOne(database.Ctx, updateDoc, filter, option)
	if e != nil {
		return "", errors.New("no  user updated")
	}
	return up.UpsertedID.(primitive.ObjectID).Hex(), nil

}

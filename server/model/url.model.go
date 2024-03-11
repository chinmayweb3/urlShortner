package model

import (
	"errors"
	"time"

	"github.com/chinmayweb3/urlshortner/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Url struct {
	LUrl       string    `json:"lUrl,omitempty" bson:"lUrl"`
	SUrl       string    `json:"sUrl"  bson:"sUrl"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	LastViewed time.Time `json:"lastViewed" bson:"lastViewed"`
	UserIp     string    `json:"userIp" bson:"userIp"`
}

func AddUrlToDb(u Url) (string, error) {

	i, err := database.Db.Collection("shorturls").InsertOne(database.Ctx, u)
	if err != nil {
		return "", errors.New("insert Failed")
	}

	return i.InsertedID.(primitive.ObjectID).Hex(), nil
}

func FindUrlBySUrl(sUrl string) (Url, error) {
	var url Url
	filter := bson.D{{Key: "sUrl", Value: sUrl}}
	if err := database.Db.Collection("shorturls").FindOne(database.Ctx, filter).Decode(&url); err != nil {
		return url, errors.New("no URL found")
	}

	return url, nil

}

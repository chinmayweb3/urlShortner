package model

import (
	"errors"
	"time"

	d "github.com/chinmayweb3/urlshortner/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Url struct {
	LUrl         string    `json:"lUrl,omitempty" bson:"lUrl"`
	SUrl         string    `json:"sUrl"  bson:"sUrl"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	LastViewed   time.Time `json:"lastViewed" bson:"lastViewed"`
	CountVisited int       `json:"countVisited" bson:"countVisited"`
	UserIp       string    `json:"userIp" bson:"userIp"`
}

func (u *Url) InsertUrl() (string, error) {

	i, err := d.Col.Surl().InsertOne(d.Ctx, u)
	if err != nil {
		return "", errors.New("insert Failed")
	}
	opts := options.Index().SetUnique(true)
	indexPro := bson.D{{Key: "sUrl", Value: 1}}
	d.Col.Surl().Indexes().CreateOne(d.Ctx, mongo.IndexModel{Keys: indexPro, Options: opts})
	// fmt.Printf("\n  adding indexing %+v or error %+v\n", str, err)

	return i.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (u *Url) FindUrlBySUrl() (Url, error) {

	var url Url
	filter := bson.D{{Key: "sUrl", Value: u.SUrl}}
	upCount := primitive.E{Key: "$inc", Value: bson.D{{Key: "countVisited", Value: 1}}}
	upViewed := primitive.E{Key: "$currentDate", Value: bson.D{{Key: "lastViewed", Value: true}}}
	update := bson.D{upCount, upViewed}

	if err := d.Col.Surl().FindOneAndUpdate(d.Ctx, filter, update).Decode(&url); err != nil {
		return url, errors.New("no URL found")
	}

	return url, nil
}

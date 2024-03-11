package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/chinmayweb3/urlshortner/database"
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
	s := i.InsertedID.(primitive.ObjectID).Hex()
	fmt.Printf("\ninserted %v type %T\n", s, s)
	return s, nil
}

package model

import (
	"errors"
	"time"

	d "github.com/chinmayweb3/urlshortner/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UserIp     string    `json:"userIp" bson:"userIp"`
	UrlLimit   int       `json:"urlLimit" bson:"urlLimit"`
	LastViewed time.Time `json:"lastViewed" bson:"lastViewed"`
}

func (u *User) FindUserByIp() (User, error) {
	var findUser User
	filter := bson.D{{Key: "userIp", Value: u.UserIp}}
	if err := d.Col.User().FindOne(d.Ctx, filter).Decode(&findUser); err != nil {
		return findUser, errors.New("no User found")
	}
	return findUser, nil
}

func (u *User) UserUpdate() (string, error) {
	updateDoc := bson.D{{Key: "userIp", Value: u.UserIp}}
	filter := bson.D{{Key: "$set", Value: u}}
	option := options.Update().SetUpsert(true)
	_, e := d.Col.User().UpdateOne(d.Ctx, updateDoc, filter, option)
	if e != nil {
		return "", errors.New("no user updated")
	}
	return "updated", nil

}

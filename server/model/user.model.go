package model

import (
	"errors"
	"time"

	"github.com/chinmayweb3/urlshortner/database"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UserIp     string    `json:"userIp" bson:"userIp"`
	UrlLimit   int       `json:"urlLimit" bson:"urlLimit"`
	LastViewed time.Time `json:"lastViewed" bson:"lastViewed"`
}

// FindUserByIp
// func (u *User) FindUserByIp() (User, error) {
// 	var findUser User
// 	if err := database.Db.Collection("users").FindOne(database.Ctx, bson.D{{Key: "userIp", Value: u.UserIp}}).Decode(&findUser); err != nil {
// 		return findUser, errors.New("no User found")
// 	}
// 	return findUser, nil
// }

func FindUserByIp(uIp string) (User, error) {
	var findUser User
	if err := database.Db.Collection("users").FindOne(database.Ctx, bson.D{{Key: "userIp", Value: uIp}}).Decode(&findUser); err != nil {
		return findUser, errors.New("no User found")
	}
	return findUser, nil
}

func UserUpdate(u User) {
	database.Db.Collection("users").UpdateOne(database.Ctx, bson.D{{Key: "userIp", Value: u.UserIp}}, u)
}

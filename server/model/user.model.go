package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/chinmayweb3/urlshortner/database"
	"go.mongodb.org/mongo-driver/bson"
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

func UserUpdate(u User) {
	filter := bson.D{{Key: "$set", Value: u}}
	up, e := database.Db.Collection("users").UpdateOne(database.Ctx, bson.D{{Key: "userIp", Value: u.UserIp}}, filter, options.Update().SetUpsert(true))
	if e != nil {
		fmt.Println("not Updated", e)
	} else {
		fmt.Printf(" Updated %+v", up)

	}

}

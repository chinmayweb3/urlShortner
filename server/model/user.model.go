package model

import "time"

type User struct {
	CreatedAt time.Duration `json:"createdAt" bson:"createdAt"`
	UserIp    string        `json:"userIp" bson:"userIp"`
	UrlLimit  int           `json:"urlLimit" bson:"urlLimit"`
}

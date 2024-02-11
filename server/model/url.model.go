package model

import "time"

type Url struct {
	LUrl       string    `json:"lUrl,omitempty" bson:"lUrl"`
	SUrl       string    `json:"sUrl"  bson:"sUrl"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	LastViewed time.Time `json:"lastViewed" bson:"lastViewed"`
	CreatedIp  string    `json:"createdIp" bson:"createdIp"`
}

package model

import "time"

type Url struct {
	LUrl       string        `json:"lUrl,omitempty" bson:"lUrl"`
	SUrl       string        `json:"sUrl"  bson:"sUrl"`
	CreatedAt  time.Duration `json:"createdAt" bson:"createdAt"`
	LastViewed time.Duration `json:"lastViewed" bson:"lastViewed"`
	CreatedIp  string        `json:"createdIp" bson:"createdIp"`
}

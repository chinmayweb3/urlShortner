package model

type TUrl struct {
	OgUrl   string `bson:"first_namW"`
	TinyUrl string `bson:"tinyUrl,omitempty"`
}

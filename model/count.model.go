package model

import (
	"log"

	d "github.com/chinmayweb3/urlshortner/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Count struct {
	Name string `json:"name,omitempty" bson:"name"`
	Num  int64  `json:"num,omitempty" bson:"num"`
}

func CreateNumber() {

	filter := bson.D{{Key: "name", Value: "num"}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "num", Value: 0}}}}

	opts := options.Update().SetUpsert(true)
	_, err := d.Col.Count().UpdateOne(d.Ctx, filter, update, opts)
	if err != nil {
		log.Panicln("Count Failed: ", err)
	}
}

package model

import (
	"log"

	"github.com/chinmayweb3/urlshortner/database"
	d "github.com/chinmayweb3/urlshortner/database"
	"go.mongodb.org/mongo-driver/bson"
)

type Count struct {
	Name string `json:"name,omitempty" bson:"name"`
	Num  int64  `json:"num,omitempty" bson:"num"`
}

func CreateCountNumber() {
	find := bson.D{{Key: "name", Value: "num"}}
	add := bson.D{{Key: "name", Value: "num"}, {Key: "num", Value: 0}}

	f := d.Col.Count().FindOne(d.Ctx, find)
	if f.Err() != nil {
		d.Col.Count().InsertOne(d.Ctx, add)
	}
}

func GetCountNumber() int64 {
	var c Count
	filter := bson.D{{Key: "name", Value: "num"}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "num", Value: 1}}}}

	err := d.Col.Count().FindOneAndUpdate(database.Ctx, filter, update).Decode(&c)
	// fmt.Println("this is the number ", err, c)
	if err != nil {
		log.Panicln("Get count number error: ", err)
	}
	return c.Num
}

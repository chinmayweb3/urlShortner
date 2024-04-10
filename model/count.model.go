package model

import (
	"log"

	"github.com/chinmayweb3/urlshortner/database"
	d "github.com/chinmayweb3/urlshortner/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Count struct {
	Name string `json:"name,omitempty" bson:"name"`
	Num  int64  `json:"num,omitempty" bson:"num"`
}

func CreateCountNumber() {

	filter := bson.D{{Key: "name", Value: "num"}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "num", Value: 0}}}}

	opts := options.Update().SetUpsert(true)
	_, err := d.Col.Count().UpdateOne(d.Ctx, filter, update, opts)
	if err != nil {
		log.Panicln("Count Failed: ", err)
	}
}

func GetCountNumber() int64 {
	var c Count
	filter := bson.D{{"name", "num"}}
	update := bson.D{{"$inc", bson.D{{"num", 1}}}}

	err := d.Col.Count().FindOneAndUpdate(database.Ctx, filter, update).Decode(&c)
	// fmt.Println("this is the number ", err, c)
	if err != nil {
		log.Panicln("Get count number error: ", err)
	}
	return c.Num
}

package api

import (
	"fmt"
	"log"

	"github.com/chinmayweb3/urlshortner/config"
	"github.com/chinmayweb3/urlshortner/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetHandler(c *gin.Context) {
	_ = c.Param("tinyUrl")
	var a model.Url

	col := config.Database.Collection("shorturls")

	t := model.Url{SUrl: "path"}
	resp, _ := col.InsertOne(config.Ctx, t)

	log.Println("inserting a row", resp)

	err := col.FindOne(config.Ctx, bson.D{{Key: "sUrl", Value: "path"}}).Decode(&a)

	if err != nil {
		log.Println("log error ", err)
		c.JSON(404, fmt.Sprintln("error couldn't find it"))
		return
	}

	log.Println("decoding of resp", a)

	c.JSON(200, a)

}

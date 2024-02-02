package api

import (
	"fmt"
	"log"

	"github.com/chinmayweb3/urlshortner/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type ReceiveUrl struct {
	TinyUrl string `json:"tinyUrl"`
	OgUrl   string `json:"ogUrl"`
}

func GetHandler(c *gin.Context) {
	tinyUrl := c.Param("tinyUrl")

	var a ReceiveUrl

	col := config.Database.Collection("shorturls")

	err := col.FindOne(config.Ctx, bson.D{{Key: "tinyUrl", Value: tinyUrl}}).Decode(&a)
	if err != nil {
		log.Println("log error ", err)
		c.JSON(404, fmt.Sprintln("error couldn't find it"))
		return

	}

	log.Println("decoding of resp", a)

	c.JSON(200, fmt.Sprintln("this is the tinyUrl test11", a))

}

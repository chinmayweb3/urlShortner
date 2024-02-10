package api

import (
	"log"

	"github.com/chinmayweb3/urlshortner/config"
	"github.com/gin-gonic/gin"
)

func TestApi(c *gin.Context) {

	log.Println("log of client: ", c.ClientIP())
	log.Println("log of remote: ", c.RemoteIP())
	log.Println("log of L: ")
	c.JSON(200, c.ClientIP())

	config.Database.Collection("shorturls").InsertOne(config.Ctx, map[string]string{"url2": "test3"})

	// c.JSON(200, "data entered")

}

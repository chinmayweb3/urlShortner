package api

import (
	"github.com/chinmayweb3/urlshortner/config"
	"github.com/gin-gonic/gin"
)

func TestApi(c *gin.Context) {

	config.Database.Collection("shorturls").InsertOne(config.Ctx, map[string]string{"url2": "test3"})

	c.JSON(200, "data entered")

}

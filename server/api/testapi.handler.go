package api

import (
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func TestApi(c *gin.Context) {

	// log.Println("log of client: ", c.ClientIP())
	// log.Println("log of remote: ", c.RemoteIP())
	// log.Println("log of L: ")
	// c.JSON(200, c.ClientIP())
	// r := c.Request.Header.Get("X-FORWARDED-FOR")
	c.PostForm("url")
	var reqUrl struct {
		Url string `json:"url"`
	}

	e, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(e, &reqUrl)

	log.Println("req", reqUrl)

	b, _ := io.ReadAll(c.Request.Body)

	// log.Println("x-", c.Request.Header.Get("X-FORWARDED-FOR"))
	log.Println("x-", c.RemoteIP())
	log.Println("micro-", string(b))
	log.Println("micro-", c.PostForm("url"))
	log.Println("micro-", reqUrl.Url == "")

	// database.Db.Collection("shorturls").InsertOne(database.Ctx, map[string]string{"url2": "test3"})

	c.JSON(200, gin.H{"hour": time.Now(), "micro": time.UnixDate, "value": reqUrl.Url == ""})

}

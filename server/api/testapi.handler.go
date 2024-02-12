package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"time"

	"github.com/chinmayweb3/urlshortner/helper"
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
	u, err := url.ParseRequestURI(reqUrl.Url)

	log.Println("encoding 62", helper.Base62Encoding(time.Now()))

	log.Println("url parse ", u)
	log.Println("url parse error ", err)

	b, _ := io.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()

	// log.Println("x-", c.Request.Header.Get("X-FORWARDED-FOR"))
	log.Println("x-", c.RemoteIP())
	log.Println("micro-", string(b))
	log.Println("micro-", c.PostForm("url"))
	log.Println("micro-", reqUrl.Url == "")

	// database.Db.Collection("shorturls").InsertOne(database.Ctx, map[string]string{"url2": "test3"})

	fmt.Println(time.Now().Format("2006-01-0215:04:05:10"))
	c.JSON(200, gin.H{"hour": time.Now(), "sdf": time.Now().Format("06-01-02 15:04:05"), "micro": time.UnixDate, "value": helper.Base62Encoding(time.Now()), "parse url": u})

}

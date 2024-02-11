package api

import (
	"encoding/json"
	"io"
	"time"

	"github.com/chinmayweb3/urlshortner/model"
	"github.com/gin-gonic/gin"
)

func Shortener(c *gin.Context) {

	user := model.User{
		UserIp: c.ClientIP(),
	}
	var reqUrl struct {
		Url string `json:"url"`
	}

	e, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(e, &reqUrl)

	//check if they only pass the json value as key url
	if reqUrl.Url == "" {
		c.JSON(429, gin.H{"error": "URL not Found."})
		return
	}

	// Check if the url is the correct url

	// 	Check if the user for url limit within the database
	findUser, err := user.FindUserByIp()
	// If the user is not present then create one
	if err != nil {
		user.CreatedAt = time.Now()
	}
	// If url limit is 0 then return error for exhaust url limit
	if findUser.UrlLimit == 0 {
		c.JSON(429, gin.H{"error": "URL Limit Exhausted."})
		return
	}
	// if url is not correct then return error for incorrect url
	// If all the things are true start the process for encoding the url
	// Encode the current date and time to base62
	// Take the first 7 character in a variable
	// Add the domain name to the start of the 7 character
	// Take the url and respected struct and put it in the database
	// Return the url back to the user
	c.JSON(200, "working post")
}

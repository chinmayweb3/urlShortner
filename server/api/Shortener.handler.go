package api

import (
	"encoding/json"
	"io"
	"net/url"
	"time"

	"github.com/chinmayweb3/urlshortner/helper"
	"github.com/chinmayweb3/urlshortner/model"
	"github.com/gin-gonic/gin"
)

func Shortener(c *gin.Context) {

	currentTime := time.Now()

	user := model.User{
		UserIp:     c.ClientIP(),
		UrlLimit:   10,
		CreatedAt:  currentTime,
		LastViewed: currentTime,
	}

	var reqUrl struct {
		Url string `json:"url"`
	}
	// json.Unmarshal the api request body in req Url
	e, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(e, &reqUrl)

	//check if they only pass the json value as key url
	if reqUrl.Url == "" {
		c.JSON(429, gin.H{"error": "URL not Found."})
		return
	}

	// Check if the url is the correct url
	_, err := url.ParseRequestURI(reqUrl.Url)
	if err != nil {
		c.JSON(200, err)
	}

	// 	Check if the user for url limit within the database
	findUser, err := user.FindUserByIp()

	// If the user is not present then create a new user
	if err == nil {
		user.CreatedAt = findUser.CreatedAt
		user.UrlLimit = findUser.UrlLimit - 1 // Reduce the count of urls by one from the existing user's url_limit

	}

	// If url limit is 0 then return error for exhaust url limit
	if user.UrlLimit < 0 {
		c.JSON(429, gin.H{"error": "URL Limit Exhausted."})
		return
	}

	// If all the things are true start the process for encoding the url
	// Encode the current date and time to base62
	_ = helper.Base62Encoding(currentTime)

	// Take the first 7 character in a variable
	// Add the domain name to the start of the 7 character
	// Take the url and respected struct and put it in the database
	// Return the url back to the user
	c.JSON(200, "working post")
}

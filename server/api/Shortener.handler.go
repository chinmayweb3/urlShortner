package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/chinmayweb3/urlshortner/helper"
	"github.com/chinmayweb3/urlshortner/model"
	"github.com/gin-gonic/gin"
)

func Shortener(c *gin.Context) {

	currentTime := time.Now()

	user := model.User{
		UserIp:     c.ClientIP(),
		UrlLimit:   9,
		CreatedAt:  currentTime,
		LastViewed: currentTime,
	}

	var reqUrl struct {
		Url string `json:"url"`
	}
	// json.Unmarshal the api request body in req Url
	e, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(e, &reqUrl)

	defer c.Request.Body.Close()

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
	findUser, err := model.FindUserByIp(user.UserIp)

	// If the user is not present then create a new user
	if err == nil {
		user.CreatedAt = findUser.CreatedAt
		user.LastViewed = currentTime
		user.UrlLimit = findUser.UrlLimit - 1 // Reduce the count of urls by one from the existing user's url_limit

	}

	// If url limit is 0 then return error for exhaust url limit
	if user.UrlLimit < 0 {
		c.JSON(429, gin.H{"error": "URL Limit Exhausted."})
		return
	}

	// If all the things are true start the process for encoding the url
	// Encode the helper.base62 Take the first 12 character in a variable
	b62 := helper.Base62Encoding()

	// Add the domain name to the start of the 12 character
	domain := c.Request.Host
	var SUrl string
	if strings.Contains(domain, "localhost") || strings.Contains(domain, "127.0.0.") {
		SUrl = fmt.Sprintf("http://%v/%v", domain, b62)
	} else {
		SUrl = fmt.Sprintf("https://%v/%v", domain, b62)
	}

	// Take the url and respected struct and put it in the database
	encodeUrl := model.Url{
		LUrl:       reqUrl.Url,
		SUrl:       SUrl,
		CreatedAt:  currentTime,
		LastViewed: currentTime,
		UserIp:     c.ClientIP(),
	}

	// modify the database in users collection

	// add url in the database in urls collection
	model.AddUrlToDb(encodeUrl)

	// Return the url back to the user
	c.JSON(200, encodeUrl)
}

package api

import "github.com/gin-gonic/gin"

func Shortener(c *gin.Context) {

	// userIp := c.ClientIP()

	// 	Check if the user for url limit within the database

	// If url limit is 0 then return error for exhaust url limit
	// If the user is not present then create one
	// Check if the url is the correct url
	// if url is not correct then return error for incorrect url
	// If all the things are true start the process for encoding the url
	// Encode the current date and time to base62
	// Take the first 7 character in a variable
	// Add the domain name to the start of the 7 character
	// Take the url and respected struct and put it in the database
	// Return the url back to the user
	c.JSON(200, "working post")
}

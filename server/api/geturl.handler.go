package api

import (
	"github.com/chinmayweb3/urlshortner/model"
	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	s := c.Param("shortUrl")

	//find the short url in database
	u, err := model.FindUrlBySUrl(s)
	if err != nil {
		c.JSON(200, gin.H{"error": err})
	}

	c.Redirect(200, u.LUrl)
}

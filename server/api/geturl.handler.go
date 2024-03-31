package api

import (
	"github.com/chinmayweb3/urlshortner/model"
	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {
	s := c.Param("shortUrl")

	url, err := model.FindUrlBySUrl(s)
	if err != nil {
		c.JSON(404, gin.H{"error": "URL not exist"})
		return
	}

	c.Redirect(301, url.LUrl)
}

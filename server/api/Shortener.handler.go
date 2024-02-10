package api

import "github.com/gin-gonic/gin"

func Shortener(c *gin.Context) {
	c.JSON(200, "working post")
}

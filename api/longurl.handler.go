package api

import "github.com/gin-gonic/gin"

func Longurl(c *gin.Context) {
	c.JSON(200, "working post")
}

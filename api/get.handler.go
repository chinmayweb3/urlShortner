package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context) {

	tinyUrl := c.Param("tinyUrl")

	// fmt.Printf("GET %s", tinyUrl)

	c.JSON(200, fmt.Sprintln("this is the tinyUrl", tinyUrl))

}

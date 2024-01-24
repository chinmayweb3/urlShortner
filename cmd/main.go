package main

import (
	"fmt"
	"os"

	"github.com/chinmayweb3/urlshortner/api"
	"github.com/gin-gonic/gin"
)

func main() {

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	r := gin.Default()

	r.GET("/:tinyUrl", api.GetHandler) //pending
	r.POST("/long", api.Longurl)       //pending

	if err := r.Run(":" + PORT); err != nil {
		fmt.Println("there is a problem in routing ", err)
	}

}

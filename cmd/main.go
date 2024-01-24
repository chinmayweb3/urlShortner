package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chinmayweb3/urlshortner/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	PORT := os.Getenv("PORT")

	r := gin.Default()

	r.GET("/:tinyUrl", api.GetHandler) //pending
	r.POST("/long", api.Longurl)       //pending

	if err := r.Run(":" + PORT); err != nil {
		fmt.Println("there is a problem in routing ", err)
	}

}

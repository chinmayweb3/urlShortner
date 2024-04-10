package main

import (
	"log"
	"os"

	"github.com/chinmayweb3/urlshortner/api"
	"github.com/chinmayweb3/urlshortner/database"
	"github.com/chinmayweb3/urlshortner/model"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Initialize()     // Initialize the database connection
	model.CreateCountNumber() // Create a new number for each init
	model.UrlInitialize()
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file:", err)
	// }

	PORT := os.Getenv("PORT")
	r := gin.Default()

	// r.GET("/", api.TestApi)
	// r.POST("/", api.TestApi)
	r.GET("/:shortUrl", api.GetHandler)        //completed
	r.POST("/api/v1/shortener", api.Shortener) //completed

	if err := r.Run(":" + PORT); err != nil {
		log.Panicln("routing Failed : ", err)
	}

}

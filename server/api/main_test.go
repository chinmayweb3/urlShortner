package api

import (
	"fmt"
	"os"
	"testing"

	"github.com/chinmayweb3/urlshortner/database"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	database.Initialize() // Initialize the database connection

	PORT := ":3001"
	r := gin.Default()

	r.GET("/", TestApi)                    //pending
	r.POST("/", TestApi)                   //pending
	r.GET("/:shortUrl", GetHandler)        //pending
	r.POST("/api/v1/shortener", Shortener) //complete

	if err := r.Run(PORT); err != nil {
		fmt.Println("there is a problem in routing ", err)
	}
	os.Exit(m.Run())

}

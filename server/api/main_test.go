package api

import (
	"os"
	"testing"

	"github.com/chinmayweb3/urlshortner/database"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	database.TestInit() // Initialize the database connection
	os.Exit(m.Run())

}

var TestHostname = "http://localhost:3001/"

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

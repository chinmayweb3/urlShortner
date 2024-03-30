package api

import (
	"os"
	"testing"

	"github.com/chinmayweb3/urlshortner/database"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	database.Initialize() // Initialize the database connection
	os.Exit(m.Run())

}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

package api

import (
	"os"
	"testing"

	"github.com/chinmayweb3/urlshortner/database"
)

func TestMain(m *testing.M) {
	database.TestInit() // Initialize the database connection
	os.Exit(m.Run())

}

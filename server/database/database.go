package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

var Ctx = context.Background()

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	var uri = os.Getenv("MONGODB_URL")

	//  Create a client to the MongoDB server.
	log.Println("the database is connecting... ")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(Ctx, opts)
	if err != nil {
		log.Println("Failed to connect with the DB ", err, "\n", uri)
		panic(err)
	}

	// defer func() {
	// 	client.Disconnect(Ctx)
	// }()

	Db = client.Database("urlshortner")
	log.Println("database has been established")

}

//Not using it for now ----------
//Not using it for now ----------

// func CreateClient() (*mongo.Database, error) {

// 	var uri = os.Getenv("MONGODB_URL")
// 	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
// 	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

// 	client, err := mongo.Connect(Ctx, opts)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client.Database("urlshortner"), nil

// }

// func DeferClient(c *mongo.Database) error {
// 	return c.Client().Disconnect(Ctx)
// }

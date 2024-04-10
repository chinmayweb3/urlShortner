package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

var Ctx = context.Background()

func Initialize() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file:", err)
	// }

	var uri = os.Getenv("MONGODB_URL")

	//  Create a client to the MongoDB server.
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(Ctx, opts)
	if err != nil {
		log.Println("Failed to connect with the DB ", err, "\n", uri)
		panic(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Println("Ping failed to connect with the DB ", err, "\n", uri)
		panic(err)
	}

	Db = client.Database("urlshortner")
	log.Println("database has been established")

}

func TestInit() {
	var uri = "mongodb://localhost:27017"

	//  Create a client to the MongoDB server.
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(Ctx, opts)
	if err != nil {
		log.Println("Failed to connect with the DB ", err, "\n", uri)
		panic(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Println("Ping failed to connect with the DB ", err, "\n", uri)
		panic(err)
	}

	Db = client.Database("urlshortner")
	log.Println("database has been established")

}

type iCols struct{}

func (c *iCols) User() *mongo.Collection {
	return Db.Collection("users")
}
func (c *iCols) Surl() *mongo.Collection {
	return Db.Collection("shorturls")
}

func (c *iCols) Count() *mongo.Collection {
	return Db.Collection("count")
}

var Col = &iCols{}

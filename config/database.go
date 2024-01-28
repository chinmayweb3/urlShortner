package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	var uri = os.Getenv("MONGODB_URL")

	//  Create a client to the MongoDB server.
	log.Println("the database is connecting... ")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	Database = client.Database("urlshortner")
	log.Println("database has been established")

	// Database.Collection("shorturls").InsertOne(context.TODO(), map[string]string{"testname": "testvalue"})

}

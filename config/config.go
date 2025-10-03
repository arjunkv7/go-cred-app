package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbClient *mongo.Client

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	DbClient = client
	return DbClient
}

func EnvMongoURI() string {
    return os.Getenv("MONGO_URI")
}
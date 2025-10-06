package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DbClient *mongo.Database

func ConnectDB() *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	DbClient = client.Database("crud-app")
	return DbClient
}

func EnvMongoURI() string {
    return os.Getenv("MONGO_URI")
}
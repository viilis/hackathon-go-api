package db

import (
	"context"
	"log"

	"github.com/viilis/go-api/app/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DbRef *mongo.Database
	MONGO_URI = "mongodb://localhost:27017" //dev
)

func InitDb() {
	log.Println("Starting database")
	
	clientOptions := options.Client().ApplyURI(utils.Config.DbUri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")

	DbRef = client.Database("go-api-db")
}
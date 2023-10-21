package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Collection

// type DatabaseRefs struct {
// 	dbRef *mongo.Collection
// }

func InitDb() {
	log.Println("Starting database")
	
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected, getting collectio references")

	Database = findCollections(client)
}

func findCollections(client *mongo.Client) (*mongo.Collection) {
	return client.Database("go-api-db").Collection("todos")
}
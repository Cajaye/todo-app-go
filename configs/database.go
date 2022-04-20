package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI()))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	err = client.Connect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
	return client
}

var DB *mongo.Client = ConnectDB()

func Collection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("golang-api").Collection(collectionName)
}

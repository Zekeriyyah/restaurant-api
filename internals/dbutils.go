package internals

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() (*mongo.Client, error) {
	mongoDb := "mongodb://localhost:27017"

	fmt.Println(mongoDb)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDb))
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

var Client, _ = DBInstance()

func OpenCollection(client *mongo.Client, nameOfCollection string) *mongo.Collection {
	collection := client.Database("restaurant").Collection(nameOfCollection)

	return collection
}

package connection

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func MongoConnection() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoServer, _ := os.LookupEnv("MONGO-SERVER")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoServer))

	if err != nil {
		log.Fatal("Error connection mongodb")
	}

	error := client.Ping(context.TODO(), nil)
	if error != nil {
		log.Fatal(error)
	}

	return client
}

func PersonCollection() *mongo.Collection {
	client := MongoConnection()

	collection := client.Database("Person").Collection("person")

	return collection
}

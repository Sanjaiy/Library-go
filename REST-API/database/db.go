package database

import (
	"context"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func StartMongoDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return errors.New("MONGODB_URI is not set")
	}
	opts := options.Client().ApplyURI(uri)

	var err error

	MongoClient, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}

	// Send a ping to confirm a successful connection
	if err := MongoClient.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return err
	}
	log.Println("Pinged your deployment. You successfully connected to MongoDB!")
	
	return nil
}

func CloseMongoDB() {
	err := MongoClient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}

func GetCollection(name string) *mongo.Collection {
	return MongoClient.Database("SK").Collection(name)
}
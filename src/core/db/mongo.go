package db

import (
	"context"
	"fmt"
	"go-boilerplate/src/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Database

func connectMongoDb(connectionString string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(config.LoadConfig(("MONGO_DB")))

	fmt.Println("Connected to MongoDB!")

	return db, nil
}

func InitMongoDB() *mongo.Database {
	var error error
	client, error = connectMongoDb(config.LoadConfig(("MONGO_CONNECTION_STRING")))

	if error != nil {
		log.Fatal(error)
	}

	return client
}

func GetMongoDb() *mongo.Database {
	return client
}

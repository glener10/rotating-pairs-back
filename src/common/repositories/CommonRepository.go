package CommonRepository

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(collectionName string) (*mongo.Collection, error) {
	dataBaseName := os.Getenv("MONGODB_DATABASE_NAME")
	if dataBaseName == "" {
		return nil, errors.New("MONGODB_DATABASE_NAME is not defined")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		return nil, errors.New("MONGODB_URI is not defined")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	coll := client.Database(dataBaseName).Collection(collectionName)
	return coll, nil
}

func Disconnect(collection *mongo.Collection) {
	client := collection.Database().Client()
	err := client.Disconnect(context.TODO())
	if err != nil {
		fmt.Println("Error to disconnect database.")
	}
}

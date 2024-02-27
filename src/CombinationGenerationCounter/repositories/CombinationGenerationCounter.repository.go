package CombinationGenerationCounterRepo

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Collection, error) {

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		return nil, errors.New("MONGODB_URI is not defined")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	dataBaseName := os.Getenv("MONGODB_DATABASE_NAME")
	if dataBaseName == "" {
		return nil, errors.New("MONGODB_DATABASE_NAME is not defined")
	}

	coll := client.Database(dataBaseName).Collection("TotalCombinationGenerationAccordingNumberOfEntries")

	if coll == nil {
		return nil, errors.New("Error to connect database!")
	}

	return coll, nil
}

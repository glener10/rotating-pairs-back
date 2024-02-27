package CombinationGenerationCounterRepo

import (
	"context"
	"errors"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Collection, error) {
	err := godotenv.Load("../../../")
	if err != nil {
		return nil, errors.New("Erro to load environment variables")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		return nil, errors.New("MONGODB_URI is not defined")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	coll := client.Database("test").Collection("TotalCombinationGenerationAccordingNumberOfEntries")

	return coll, nil
}

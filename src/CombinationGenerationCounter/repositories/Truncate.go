package CombinationGenerationCounterRepo

import (
	"context"
	"errors"

	CommonRepository "github.com/glener10/rotating-pairs-back/src/common/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

func Truncate() error {
	if err := CleanCollection(); err != nil {
		return errors.New(err.Error())
	}
	if err := PopulateCollection(); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func CleanCollection() error {
	collectionName := "TotalCombinationGenerationAccordingNumberOfEntries"
	col, err := CommonRepository.Connect(collectionName)
	if err != nil {
		return errors.New(err.Error())
	}
	_, err = col.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return errors.New("Error to DeleteMany in 'TotalCombinationGenerationAccordingNumberOfEntries' collection:" + err.Error())
	}
	return nil
}

func PopulateCollection() error {
	collectionName := "TotalCombinationGenerationAccordingNumberOfEntries"
	col, err := CommonRepository.Connect(collectionName)
	if err != nil {
		return errors.New("failed to connect to database: " + err.Error())
	}

	one := int32(1)
	documents := []bson.M{
		{
			"NumberOfEntries": 1,
			"Count":           one,
		},
		{
			"NumberOfEntries": 2,
			"Count":           one,
		},
	}

	for _, doc := range documents {
		_, err := col.InsertOne(context.Background(), doc)
		if err != nil {
			return errors.New("failed to insert document: " + err.Error())
		}
	}
	return nil
}

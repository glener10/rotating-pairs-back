package CombinationUtils

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
	collectionName := "Combinations"
	col, err := CommonRepository.Connect(collectionName)
	if err != nil {
		return errors.New(err.Error())
	}
	_, err = col.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return errors.New("Error to DeleteMany in 'Combinations' collection:" + err.Error())
	}
	return nil
}

func PopulateCollection() error {
	collectionName := "Combinations"
	col, err := CommonRepository.Connect(collectionName)
	if err != nil {
		return errors.New("failed to connect to database: " + err.Error())
	}
	documents := []bson.M{
		{
			"NumberOfInputs":                2,
			"NumberOfSprints":               1,
			"NumberOfCombinationsPerSprint": 1,
			"Sprints": []bson.M{
				{
					"Combinations": []bson.M{
						{
							"PairOne": 0,
							"PairTwo": 1,
						},
					},
				},
			},
		},
		{
			"NumberOfInputs":                3,
			"NumberOfSprints":               3,
			"NumberOfCombinationsPerSprint": 2,
			"Sprints": []bson.M{
				{
					"Combinations": []bson.M{
						{
							"PairOne": 0,
							"PairTwo": 1,
						},
						{
							"PairOne": 2,
							"PairTwo": 2,
						},
					},
				},
				{
					"Combinations": []bson.M{
						{
							"PairOne": 0,
							"PairTwo": 2,
						},
						{
							"PairOne": 1,
							"PairTwo": 1,
						},
					},
				},
				{
					"Combinations": []bson.M{
						{
							"PairOne": 1,
							"PairTwo": 2,
						},
						{
							"PairOne": 0,
							"PairTwo": 0,
						},
					},
				},
			},
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

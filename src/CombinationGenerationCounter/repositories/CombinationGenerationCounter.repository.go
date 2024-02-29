package CombinationGenerationCounterRepo

import (
	"context"
	"errors"
	"strconv"

	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	CommonRepository "github.com/glener10/rotating-pairs-back/src/common/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindByNumberOfInputs(NumberOfInputs int16) (*CombinationGenerationCounterEntity.CombinationGenerationCounter, error) {
	collectionName := "TotalCombinationGenerationAccordingNumberOfInputs"
	col, err := CommonRepository.Connect(collectionName)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer CommonRepository.Disconnect(col)

	ctx := context.TODO()
	filter := bson.D{{Key: "NumberOfInputs", Value: NumberOfInputs}}
	var result CombinationGenerationCounterEntity.CombinationGenerationCounter
	err = col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, errors.New("Error to find a Combination Generation Counter with " + strconv.Itoa(int(NumberOfInputs)) + " inputs")
	}
	return &result, nil
}

func CreateCombinationGenerationCounter(NumberOfInputs int16) (*CombinationGenerationCounterEntity.CombinationGenerationCounter, error) {
	collectionName := "TotalCombinationGenerationAccordingNumberOfInputs"
	col, err := CommonRepository.Connect(collectionName)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer CommonRepository.Disconnect(col)

	ctx := context.TODO()
	document := bson.M{
		"NumberOfInputs": NumberOfInputs,
		"Count":          1,
	}
	_, err = col.InsertOne(ctx, document)
	if err != nil {
		return nil, errors.New("Error to insert new Combination Generation Counter with " + strconv.Itoa(int(NumberOfInputs)) + " inputs")
	}
	return &CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfInputs: NumberOfInputs,
		Count:          1,
	}, nil
}

func IncrementCombinationGenerationCounter(combination *CombinationGenerationCounterEntity.CombinationGenerationCounter) (*CombinationGenerationCounterEntity.CombinationGenerationCounter, error) {
	collectionName := "TotalCombinationGenerationAccordingNumberOfInputs"
	col, err := CommonRepository.Connect(collectionName)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer CommonRepository.Disconnect(col)

	ctx := context.TODO()
	newCountOfCombination := combination.Count + 1
	filter := bson.D{{Key: "NumberOfInputs", Value: combination.NumberOfInputs}}
	combination.Count = newCountOfCombination
	update := bson.M{"$set": bson.M{"Count": newCountOfCombination}}
	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, errors.New("Error to increment the Combination Generation Counter with " + strconv.Itoa(int(combination.NumberOfInputs)) + " inputs")
	}

	return combination, nil
}

func ListAllCombinationsCounters() (*[]CombinationGenerationCounterEntity.CombinationGenerationCounter, error) {
	collectionName := "TotalCombinationGenerationAccordingNumberOfInputs"
	col, err := CommonRepository.Connect(collectionName)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer CommonRepository.Disconnect(col)

	ctx := context.TODO()
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "Count", Value: -1}})

	cursor, err := col.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, errors.New("failed to execute find operation: " + err.Error())
	}
	defer cursor.Close(ctx)

	var results []CombinationGenerationCounterEntity.CombinationGenerationCounter
	for cursor.Next(ctx) {
		var result CombinationGenerationCounterEntity.CombinationGenerationCounter
		if err := cursor.Decode(&result); err != nil {
			return nil, errors.New("Failed to decode document: " + err.Error())
		}
		results = append(results, result)
	}

	return &results, nil
}

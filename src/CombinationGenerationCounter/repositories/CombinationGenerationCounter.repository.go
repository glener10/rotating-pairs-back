package CombinationGenerationCounterRepo

import (
	"context"
	"errors"
	"strconv"

	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	CommonRepository "github.com/glener10/rotating-pairs-back/src/common/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

func FindByNumberOfEntries(numberOfEntries int16) (*CombinationGenerationCounterEntity.CombinationGenerationCounter, error) {
	collectionName := "TotalCombinationGenerationAccordingNumberOfEntries"
	col, _ := CommonRepository.Connect(collectionName)
	defer CommonRepository.Disconnect(col)

	ctx := context.TODO()
	filter := bson.D{{Key: "NumberOfEntries", Value: numberOfEntries}}
	var result CombinationGenerationCounterEntity.CombinationGenerationCounter
	err := col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, errors.New("Error to find a Combination Generation Counter with " + strconv.Itoa(int(numberOfEntries)) + " inputs")
	}
	return &result, nil
}

func CreateCombinationGenerationCounter(numberOfEntries int16) (*CombinationGenerationCounterEntity.CombinationGenerationCounter, error) {
	collectionName := "TotalCombinationGenerationAccordingNumberOfEntries"
	col, _ := CommonRepository.Connect(collectionName)
	defer CommonRepository.Disconnect(col)

	one := int32(1)
	ctx := context.TODO()
	document := bson.M{
		"NumberOfEntries": numberOfEntries,
		"Count":           one,
	}
	_, err := col.InsertOne(ctx, document)
	if err != nil {
		return nil, errors.New("Error to insert new Combination Generation Counter with " + strconv.Itoa(int(numberOfEntries)) + " inputs")
	}

	return &CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfEntries: numberOfEntries,
		Count:           &one,
	}, nil
}

func IncrementCombinationGenerationCounter(combination *CombinationGenerationCounterEntity.CombinationGenerationCounter) (*CombinationGenerationCounterEntity.CombinationGenerationCounter, error) {
	collectionName := "TotalCombinationGenerationAccordingNumberOfEntries"
	col, _ := CommonRepository.Connect(collectionName)
	defer CommonRepository.Disconnect(col)

	ctx := context.TODO()
	newCountOfCombination := *combination.Count + 1
	filter := bson.D{{Key: "NumberOfEntries", Value: combination.NumberOfEntries}}
	combination.Count = &newCountOfCombination
	update := bson.M{"$set": bson.M{"Count": newCountOfCombination}}
	_, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, errors.New("Error to increment the Combination Generation Counter with " + strconv.Itoa(int(combination.NumberOfEntries)) + " inputs")
	}

	return combination, nil
}

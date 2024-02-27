package CombinationGenerationCounterRepo

import (
	"context"
	"errors"
	"os"
	"strconv"

	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	CommonRepository "github.com/glener10/rotating-pairs-back/src/common/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindByNumberOfEntries(numberOfEntries int16) (*CombinationGenerationCounterEntity.CombinationGenerationCounter, error) {
	_, err := ValidateNumberOfEntries(numberOfEntries)
	if err != nil {
		return nil, err
	}

	dataBaseName := os.Getenv("MONGODB_DATABASE_NAME")
	if dataBaseName == "" {
		return nil, errors.New("MONGODB_DATABASE_NAME is not defined")
	}
	col, _ := CommonRepository.Connect(dataBaseName)
	defer CommonRepository.Disconnect(col)
	ctx := context.TODO()
	filter := bson.D{{Key: "NumberOfEntries", Value: numberOfEntries}}
	var result CombinationGenerationCounterEntity.CombinationGenerationCounter
	err = col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return CreateCombinationGenerationCounter(numberOfEntries, col)
	}
	return UpdateCombinationGenerationCounter(&result, col)
}

func CreateCombinationGenerationCounter(numberOfEntries int16, col *mongo.Collection) (*CombinationGenerationCounterEntity.CombinationGenerationCounter, error) {
	zero := int32(0)
	ctx := context.TODO()
	document := bson.M{
		"NumberOfEntries": numberOfEntries,
		"Count":           zero,
	}
	_, err := col.InsertOne(ctx, document)
	if err != nil {
		return nil, errors.New("Error to insert new Combination Generation Counter with " + strconv.Itoa(int(numberOfEntries)) + " inputs")
	}

	return &CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfEntries: numberOfEntries,
		Count:           &zero,
	}, nil
}

func ValidateNumberOfEntries(numberOfEntries int16) (*int16, error) {
	if numberOfEntries < 1 || numberOfEntries > 20 {
		return nil, errors.New("NumberOfEntries is more than 20 or less than 1")
	}
	return &numberOfEntries, nil
}

func UpdateCombinationGenerationCounter(combination *CombinationGenerationCounterEntity.CombinationGenerationCounter, col *mongo.Collection) (*CombinationGenerationCounterEntity.CombinationGenerationCounter, error) {
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

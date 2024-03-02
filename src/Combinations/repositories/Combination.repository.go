package CombinationRepo

import (
	"context"
	"errors"
	"strconv"

	CombinationEntity "github.com/glener10/rotating-pairs-back/src/Combinations/entities"
	CommonRepository "github.com/glener10/rotating-pairs-back/src/common/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

func FindCombination(NumberOfInputs int16) (*CombinationEntity.Combination, error) {
	collectionName := "Combinations"
	col, err := CommonRepository.Connect(collectionName)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer CommonRepository.Disconnect(col)

	ctx := context.TODO()
	filter := bson.D{{Key: "NumberOfInputs", Value: NumberOfInputs}}
	var result CombinationEntity.Combination
	err = col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, errors.New("Error to find a Combination with " + strconv.Itoa(int(NumberOfInputs)) + " inputs")
	}
	return &result, nil
}

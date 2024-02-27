package CombinationGenerationCounterRepo

import (
	"context"
	"errors"
	"os"

	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	CommonRepository "github.com/glener10/rotating-pairs-back/src/common/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

func FindByNumberOfEntries(numberOfEntries int16) (*CombinationGenerationCounterEntity.CombinationGenerationCounter, error) {
	dataBaseName := os.Getenv("MONGODB_DATABASE_NAME")
	if dataBaseName == "" {
		return nil, errors.New("MONGODB_DATABASE_NAME is not defined")
	}
	col, _ := CommonRepository.Connect(dataBaseName)
	defer CommonRepository.Disconnect(col)
	ctx := context.TODO()
	filter := bson.D{{Key: "NumberOfEntries", Value: numberOfEntries}}
	var result CombinationGenerationCounterEntity.CombinationGenerationCounter
	err := col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

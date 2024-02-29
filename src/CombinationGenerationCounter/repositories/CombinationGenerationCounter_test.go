package CombinationGenerationCounterRepo

import (
	"log"
	"os"
	"testing"

	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	CombinationGenerationCounterUtils "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/utils"
	Utils "github.com/glener10/rotating-pairs-back/src/common/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	if err := Utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("Error to load environment variables: %s", err.Error())
	}
	if err := CombinationGenerationCounterUtils.Truncate(); err != nil {
		log.Fatalf("Error to exec truncate method before repository tests execution: %s", err.Error())
	}
	exitCode := m.Run()
	if err := CombinationGenerationCounterUtils.CleanCollection(); err != nil {
		log.Fatalf("Error to exec cleaning collection after repository tests execution: %s", err.Error())
	}
	os.Exit(exitCode)
}

func TestCreateCombinationGenerationCounter(t *testing.T) {
	result, _ := CreateCombinationGenerationCounter(3)

	one := int32(1)
	expectedObject := &CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfInputs: 3,
		Count:          &one,
	}

	assert.Equal(t, result, expectedObject, "The create object need to be equal to expected object")
}

func TestFindByNumberOfInputs(t *testing.T) {
	result, _ := FindByNumberOfInputs(2)

	one := int32(1)
	expectedObject := &CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfInputs: 2,
		Count:          &one,
	}

	assert.Equal(t, result, expectedObject, "The object finded in MongoDB needs to be the expected object")
}

func TestFindByNumberOfInputsInCaseDoentsFindTheCombination(t *testing.T) {
	_, err := FindByNumberOfInputs(5)

	assert.Equal(t, err.Error(), "Error to find a Combination Generation Counter with 5 inputs", "The error message need to be equal to expected error message")
}

func TestIncrementCombinationGenerationCounter(t *testing.T) {
	result, _ := FindByNumberOfInputs(1)
	_, _ = IncrementCombinationGenerationCounter(result)
	result, _ = FindByNumberOfInputs(1)

	two := int32(2)
	expectedObject := &CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfInputs: 1,
		Count:          &two,
	}

	assert.Equal(t, result, expectedObject, "The count of combination generation with '1' NumberOfInputs should be incremented to 2")
}

func TestListAllCombinationsCounters(t *testing.T) {
	if err := CombinationGenerationCounterUtils.CleanCollection(); err != nil {
		log.Fatalf("Error to exec cleaning collection after repository tests execution: %s", err.Error())
	}
	_, _ = CreateCombinationGenerationCounter(1)
	combinationTwo, _ := CreateCombinationGenerationCounter(2)
	_, _ = IncrementCombinationGenerationCounter(combinationTwo)
	_, _ = IncrementCombinationGenerationCounter(combinationTwo)
	three := int32(3)
	firstExpectededObject := CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfInputs: 2,
		Count:          &three,
	}
	one := int32(1)
	secondExpectededObject := CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfInputs: 1,
		Count:          &one,
	}

	result, _ := ListAllCombinationsCounters()

	assert.Equal(t, len(*result), 2, "Slice length should be equal")
	assert.Equal(t, (*result)[0], firstExpectededObject, "Element should be equal")
	assert.Equal(t, (*result)[1], secondExpectededObject, "Element should be equal")

}

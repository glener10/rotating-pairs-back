package CombinationGenerationCounterRepo

import (
	"log"
	"os"
	"testing"

	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	Utils "github.com/glener10/rotating-pairs-back/src/common/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	if err := Utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("Error to load environment variables: %s", err.Error())
	}
	if err := Truncate(); err != nil {
		log.Fatalf("Error to exec truncate method before repository tests execution: %s", err.Error())
	}
	exitCode := m.Run()
	if err := CleanCollection(); err != nil {
		log.Fatalf("Error to exec cleaning collection after repository tests execution: %s", err.Error())
	}
	os.Exit(exitCode)
}

func TestCreateCombinationGenerationCounter(t *testing.T) {
	result, _ := CreateCombinationGenerationCounter(3)

	one := int32(1)
	expectedObject := &CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfEntries: 3,
		Count:           &one,
	}

	assert.Equal(t, result, expectedObject, "The create object need to be equal to expected object")
}

func TestFindByNumberOfEntries(t *testing.T) {
	result, _ := FindByNumberOfEntries(2)

	one := int32(1)
	expectedObject := &CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfEntries: 2,
		Count:           &one,
	}

	assert.Equal(t, result, expectedObject, "The object finded in MongoDB needs to be the expected object")
}

func TestFindByNumberOfEntriesInCaseDoentsFindTheCombination(t *testing.T) {
	_, err := FindByNumberOfEntries(5)

	assert.Equal(t, err.Error(), "Error to find a Combination Generation Counter with 5 inputs", "The error message need to be equal to expected error message")
}

func TestIncrementCombinationGenerationCounter(t *testing.T) {
	result, _ := FindByNumberOfEntries(1)
	_, _ = IncrementCombinationGenerationCounter(result)
	result, _ = FindByNumberOfEntries(1)

	two := int32(2)
	expectedObject := &CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfEntries: 1,
		Count:           &two,
	}

	assert.Equal(t, result, expectedObject, "The count of combination generation with '1' NumberOfInputs should be incremented to 2")
}

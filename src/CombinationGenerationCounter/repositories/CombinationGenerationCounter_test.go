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
	if err := Utils.LoadEnvironmentVariables("../../.env"); err != nil {
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

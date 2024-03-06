package CombinationRepo

import (
	"log"
	"os"
	"testing"

	CombinationEntity "github.com/glener10/rotating-pairs-back/src/Combinations/entities"
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

func TestFindCombination(t *testing.T) {
	expectedObject := &CombinationEntity.Combination{
		NumberOfInputs:                2,
		NumberOfSprints:               1,
		NumberOfCombinationsPerSprint: 1,
		Sprints: []CombinationEntity.Sprint{
			{
				Combinations: []CombinationEntity.Pair{
					{
						PairOne: 0,
						PairTwo: 1,
					},
				},
			},
		},
	}

	result, _ := FindCombination(2)

	assert.Equal(t, result, expectedObject, "Should return the expected object")
}

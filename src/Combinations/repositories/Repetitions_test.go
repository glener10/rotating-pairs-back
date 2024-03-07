package CombinationRepo

import (
	"testing"

	CombinationEntity "github.com/glener10/rotating-pairs-back/src/Combinations/entities"
	CombinationUtils "github.com/glener10/rotating-pairs-back/src/Combinations/utils"
	"github.com/stretchr/testify/assert"
)

func returnAllCombinations() []*CombinationEntity.Combination {
	var combinations []*CombinationEntity.Combination

	for i := 2; i <= 3; i++ {
		found, err := FindCombination(int16(i))
		if err != nil {
			continue
		}
		combinations = append(combinations, found)
	}
	return combinations
}

func TestIfThereIsARepeatedCombination(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		assert.Equal(t, CombinationUtils.CheckIfThereIsARepeatedCombination(combinations[i].Sprints), true, "Must check if there is a repeated combination in the Combinations for the mapping of NumberOfInputs")
	}
}

func TestIfAllIndexesAreValid(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		assert.Equal(t, CombinationUtils.CheckIfAllIndexesAreValid(combinations[i].Sprints, combinations[i].NumberOfInputs), true, "Must Check If All Indexes Are Valid for the mapping for the NumberOfInputs")
	}
}

func TestIfAnyInputFromThePairIsRepeatedInTheCombinationsOfASprint(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		assert.Equal(t, CombinationUtils.CheckIfAnyInputFromThePairIsRepeatedInTheCombinationsOfASprint(combinations[i].Sprints), true, "Must check if there is a repeated combination for the mapping of NumberOfInputs")
	}
}

func TestIfAllSprintsHaveAValidNumberOfCombinationsPerSprint(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		indexArrayWithNumberInputs, numberOfInputsIsOdd := CombinationUtils.ReturnArrayAndBooleanEvenOrOdd(combinations[i].NumberOfInputs)
		generateNumberOfCombinationsPerSprint := CombinationUtils.ReturnNumberOfCombinationPerSprintRoundedDown(numberOfInputsIsOdd, indexArrayWithNumberInputs)

		assert.Equal(t, generateNumberOfCombinationsPerSprint, combinations[i].NumberOfCombinationsPerSprint, "Must has a correct number of combinations per sprint")
		assert.Equal(t, CombinationUtils.CheckIfAllSprintsHaveAValidNumberOfCombinations(combinations[i].Sprints, generateNumberOfCombinationsPerSprint), true, "All elements must has a correct number of combinations per sprint")
	}
}

func TestIfAllCombinationsHaveAValidNumberOfSprint(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		indexArrayWithNumberInputs, numberOfInputsIsOdd := CombinationUtils.ReturnArrayAndBooleanEvenOrOdd(combinations[i].NumberOfInputs)
		generateNumberOfSprints := CombinationUtils.ReturnNumberOfSprints(numberOfInputsIsOdd, indexArrayWithNumberInputs)

		assert.Equal(t, generateNumberOfSprints, combinations[i].NumberOfSprints, "Must has a correct number of sprints")
		assert.Equal(t, CombinationUtils.CheckIfAllCombinationsHaveAValidNumberOfSprint(combinations[i].Sprints, generateNumberOfSprints), true, "All elements must has a correct number of sprints")
	}
}

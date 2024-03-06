package CombinationRepo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfThereIsARepeatedCombination(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		assert.Equal(t, checkIfThereIsARepeatedCombination(combinations[i].Sprints), true, "Must check if there is a repeated combination in the Combinations for the mapping of NumberOfInputs")
	}
}

func TestIfAllIndexesAreValid(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		assert.Equal(t, checkIfAllIndexesAreValid(combinations[i].Sprints, combinations[i].NumberOfInputs), true, "Must Check If All Indexes Are Valid for the mapping for the NumberOfInputs")
	}
}

func TestIfAnyInputFromThePairIsRepeatedInTheCombinationsOfASprint(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		assert.Equal(t, checkIfAnyInputFromThePairIsRepeatedInTheCombinationsOfASprint(combinations[i].Sprints), true, "Must check if there is a repeated combination for the mapping of NumberOfInputs")
	}
}

func TestIfAllSprintsHaveAValidNumberOfCombinationsPerSprint(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		indexArrayWithNumberInputs, numberOfInputsIsOdd := returnArrayAndBooleanEvenOrOdd(combinations[i].NumberOfInputs)
		generateNumberOfCombinationsPerSprint := returnNumberOfCombinationPerSprintRoundedDown(numberOfInputsIsOdd, indexArrayWithNumberInputs)

		assert.Equal(t, generateNumberOfCombinationsPerSprint, combinations[i].NumberOfCombinationsPerSprint, "Must has a correct number of combinations per sprint")
		assert.Equal(t, checkIfAllSprintsHaveAValidNumberOfCombinations(combinations[i].Sprints, generateNumberOfCombinationsPerSprint), true, "All elements must has a correct number of combinations per sprint")
	}
}

func TestIfAllCombinationsHaveAValidNumberOfSprint(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		indexArrayWithNumberInputs, numberOfInputsIsOdd := returnArrayAndBooleanEvenOrOdd(combinations[i].NumberOfInputs)
		generateNumberOfSprints := returnNumberOfSprints(numberOfInputsIsOdd, indexArrayWithNumberInputs)

		assert.Equal(t, generateNumberOfSprints, combinations[i].NumberOfSprints, "Must has a correct number of sprints")
		assert.Equal(t, checkIfAllCombinationsHaveAValidNumberOfSprint(combinations[i].Sprints, generateNumberOfSprints), true, "All elements must has a correct number of sprints")
	}
}

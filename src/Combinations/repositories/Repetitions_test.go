package CombinationRepo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfThereIsARepeatedCombination(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		assert.Equal(t, checkIfThereIsARepeatedCombination(combinations[i].Sprints), true, "Must check if there is a repeated combination in the Combinations for the mapping of NumberOfEntries")
	}
}

func TestIfAllIndexesAreValid(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		assert.Equal(t, checkIfAllIndexesAreValid(combinations[i].Sprints, combinations[i].NumberOfInputs), true, "Must Check If All Indexes Are Valid for the mapping for the NumberOfEntries")
	}
}

func TestIfAnyInputFromThePairIsRepeatedInTheCombinationsOfASprint(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		assert.Equal(t, checkIfAnyInputFromThePairIsRepeatedInTheCombinationsOfASprint(combinations[i].Sprints), true, "Must check if there is a repeated combination for the mapping of NumberOfEntries")
	}
}

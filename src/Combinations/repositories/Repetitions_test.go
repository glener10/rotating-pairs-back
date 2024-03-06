package CombinationRepo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfThereIsARepeatedCombination(t *testing.T) {
	combinations := returnAllCombinations()
	for i := 0; i < len(combinations); i++ {
		assert.Equal(t, checkIfThereIsARepeatedCombination(combinations[i].Sprints), true, "Should return the expected object")
	}
}

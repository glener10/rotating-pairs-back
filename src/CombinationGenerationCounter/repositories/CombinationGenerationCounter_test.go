package CombinationGenerationCounterRepo

import (
	"testing"

	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	Utils "github.com/glener10/rotating-pairs-back/src/common/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateCombinationGenerationCounter(t *testing.T) {
	_ = Utils.LoadEnvironmentVariables()

	result, _ := CreateCombinationGenerationCounter(3)

	one := int32(1)
	expectedObject := &CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfEntries: 3,
		Count:           &one,
	}

	assert.Equal(t, result, expectedObject, "The create object need to be equal to expected object")
}

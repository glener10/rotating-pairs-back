package CombinationRepo

import (
	"fmt"
	"math"

	CombinationEntity "github.com/glener10/rotating-pairs-back/src/Combinations/entities"
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

func returnArrayAndBooleanEvenOrOdd(numberOfInputs int) (indexArrayWithNumberInputs []string, numberOfInputsIsOdd bool) {
	indexArrayWithNumberInputs = generateIndexArrayWithSizeOfNewEntry(numberOfInputs)
	numberOfInputsIsOdd = checkIfArrayIsOdd(indexArrayWithNumberInputs)
	return indexArrayWithNumberInputs, numberOfInputsIsOdd
}

func checkIfArrayIsOdd(inputNamesInArray []string) bool {
	return len(inputNamesInArray)%2 != 0
}

func generateIndexArrayWithSizeOfNewEntry(numberOfInputs int) []string {
	var arraySizeOfInput []string
	for index := 0; index < numberOfInputs; index++ {
		arraySizeOfInput = append(arraySizeOfInput, fmt.Sprintf("%d", index))
	}
	return arraySizeOfInput
}

func returnNumberOfCombinationPerSprintRoundedDown(numberOfNamesIsOdd bool, inputNamesInArray []string) int {
	numberOfCombinationPerSprint := len(inputNamesInArray) / 2
	if numberOfNamesIsOdd {
		numberOfCombinationPerSprint++
	}
	return int(math.Floor(float64(numberOfCombinationPerSprint)))
}

func returnNumberOfSprints(numberOfNamesIsOdd bool, inputNamesInArray []string) int {
	numberOfSprints := len(inputNamesInArray) - 1
	if numberOfNamesIsOdd {
		numberOfSprints++
	}
	return numberOfSprints
}

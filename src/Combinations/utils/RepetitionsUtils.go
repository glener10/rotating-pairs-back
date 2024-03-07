package CombinationUtils

import (
	"fmt"
)

func ReturnArrayAndBooleanEvenOrOdd(numberOfInputs int) (indexArrayWithNumberInputs []string, numberOfInputsIsOdd bool) {
	indexArrayWithNumberInputs = GenerateIndexArrayWithSizeOfNewEntry(numberOfInputs)
	numberOfInputsIsOdd = CheckIfArrayIsOdd(indexArrayWithNumberInputs)
	return indexArrayWithNumberInputs, numberOfInputsIsOdd
}

func CheckIfArrayIsOdd(inputNamesInArray []string) bool {
	return len(inputNamesInArray)%2 != 0
}

func GenerateIndexArrayWithSizeOfNewEntry(numberOfInputs int) []string {
	var arraySizeOfInput []string
	for index := 0; index < numberOfInputs; index++ {
		arraySizeOfInput = append(arraySizeOfInput, fmt.Sprintf("%d", index))
	}
	return arraySizeOfInput
}

func ReturnNumberOfCombinationPerSprintRoundedDown(numberOfNamesIsOdd bool, inputNamesInArray []string) int {
	numberOfCombinationPerSprint := len(inputNamesInArray) / 2
	if numberOfNamesIsOdd {
		numberOfCombinationPerSprint++
	}
	return numberOfCombinationPerSprint
}

func ReturnNumberOfSprints(numberOfNamesIsOdd bool, inputNamesInArray []string) int {
	numberOfSprints := len(inputNamesInArray) - 1
	if numberOfNamesIsOdd {
		numberOfSprints++
	}
	return numberOfSprints
}

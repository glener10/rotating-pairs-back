package GenerateUseCases

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	CombinationRepo "github.com/glener10/rotating-pairs-back/src/Combinations/repositories"
	GenerateRequestDto "github.com/glener10/rotating-pairs-back/src/Generate/dtos"
)

func Generate(c *gin.Context, combinationRequest GenerateRequestDto.GenerateRequest) {
	NumberOfInputs := combinationRequest.NumberOfInputs
	results, err := CombinationRepo.FindCombination(NumberOfInputs)
	if err == nil {
		statusCode := http.StatusOK
		c.JSON(statusCode, gin.H{"error": "There is already a transaction with this number of entries", "statusCode": statusCode})
		return
	}
	arrayOfPositions := returnArrayOfIndexOfNumberOfInputsSize(NumberOfInputs)
	numberOfInputsIsOdd := CombinationRepo.CheckIfArrayIsOdd(arrayOfPositions)
	numberOfSprint := CombinationRepo.ReturnNumberOfSprints(numberOfInputsIsOdd, arrayOfPositions)
	numberOfCombinationPerSprint := CombinationRepo.ReturnNumberOfCombinationPerSprintRoundedDown(numberOfInputsIsOdd, arrayOfPositions)

	sizeOfLoop := returnSizeOfCombinationsLoop(arrayOfPositions, numberOfCombinationPerSprint)

	fmt.Println(sizeOfLoop, numberOfSprint)
	c.JSON(http.StatusOK, results)
}

func returnSizeOfCombinationsLoop(arrayOfPositions []string, numberOfCombinationPerSprint int) int {
	if len(arrayOfPositions) == 2 {
		return numberOfCombinationPerSprint
	}
	return numberOfCombinationPerSprint - 1
}

func returnArrayOfIndexOfNumberOfInputsSize(NumberOfInputs int16) []string {
	var array []string
	for i := 0; i < int(NumberOfInputs); i++ {
		array = append(array, strconv.Itoa(i))
	}
	return array
}

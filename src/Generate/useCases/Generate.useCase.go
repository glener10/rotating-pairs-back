package GenerateUseCases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	CombinationRepo "github.com/glener10/rotating-pairs-back/src/Combinations/repositories"
	GenerateRequestDto "github.com/glener10/rotating-pairs-back/src/Generate/interfaces"
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
	c.JSON(http.StatusOK, arrayOfPositions)
}

func returnArrayOfIndexOfNumberOfInputsSize(NumberOfInputs int16) []int {
	var array []int
	for i := 0; i < int(NumberOfInputs); i++ {
		array = append(array, i)
	}
	return array
}

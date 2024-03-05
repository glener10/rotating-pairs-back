package CombinationUseCases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	CombinationGenerationCounterRepo "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/repositories"
	CombinationRepo "github.com/glener10/rotating-pairs-back/src/Combinations/repositories"
	CombinationRequestDto "github.com/glener10/rotating-pairs-back/src/common/interfaces"
)

func Combination(c *gin.Context, combinationRequest CombinationRequestDto.CombinationRequest) {
	NumberOfInputs := combinationRequest.NumberOfInputs
	results, err := CombinationRepo.FindCombination(NumberOfInputs)
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		errorMessage := err.Error()
		c.JSON(statusCode, gin.H{"error": errorMessage, "statusCode": statusCode})
		return
	}
	_, err = CombinationGenerationCounterRepo.IncrementCombinationGenerationCounter(combinationRequest.NumberOfInputs)
	if err != nil {
		_, err = CombinationGenerationCounterRepo.CreateCombinationGenerationCounter(combinationRequest.NumberOfInputs)
		if err != nil {
			statusCode := http.StatusUnprocessableEntity
			errorMessage := err.Error()
			c.JSON(statusCode, gin.H{"error": errorMessage, "statusCode": statusCode})
			return
		}
	}
	c.JSON(http.StatusOK, results)
}

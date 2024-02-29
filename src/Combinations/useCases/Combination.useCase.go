package CombinationUseCases

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	//TODO: Need call Increment, changing the requisition body
	c.JSON(http.StatusOK, results)
}

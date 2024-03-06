package GenerateUseCases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	CombinationRepo "github.com/glener10/rotating-pairs-back/src/Combinations/repositories"
	CombinationRequestDto "github.com/glener10/rotating-pairs-back/src/common/interfaces"
)

func Generate(c *gin.Context, combinationRequest CombinationRequestDto.CombinationRequest) {
	NumberOfInputs := combinationRequest.NumberOfInputs
	results, err := CombinationRepo.FindCombination(NumberOfInputs)
	if err == nil {
		statusCode := http.StatusOK
		c.JSON(statusCode, gin.H{"error": "There is already a transaction with this number of entries", "statusCode": statusCode})
		return
	}

	c.JSON(http.StatusOK, results)
}

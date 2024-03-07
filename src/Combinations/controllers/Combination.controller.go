package CombinationController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	CombinationUseCases "github.com/glener10/rotating-pairs-back/src/Combinations/useCases"
	CombinationRequestDto "github.com/glener10/rotating-pairs-back/src/common/dtos"
)

func Combination(c *gin.Context) {
	var combinationRequest CombinationRequestDto.CombinationRequest
	if err := c.ShouldBindJSON(&combinationRequest); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "statusCode": statusCode})
		return
	}
	if err := CombinationRequestDto.Validate(&combinationRequest); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	CombinationUseCases.Combination(c, combinationRequest)
}

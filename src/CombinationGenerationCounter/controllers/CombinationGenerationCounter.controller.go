package CombinationGenerationCounterController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	CombinationGenerationCounterUseCases "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/useCases"
)

func IncrementCombinationGenerationCounter(c *gin.Context) {
	var combination CombinationGenerationCounterEntity.CombinationGenerationCounter
	if err := c.ShouldBindJSON(&combination); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "statusCode": statusCode})
		return
	}
	if err := CombinationGenerationCounterEntity.Validate(&combination); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	CombinationGenerationCounterUseCases.IncrementCombinationGenerationCounter(c, combination)
}

func ListAllCombinationsCounters(c *gin.Context) {
	CombinationGenerationCounterUseCases.ListAllCombinationsCounters(c)
}

package IncrementCombinationGenerationCounterUseCase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	CombinationGenerationCounterRepo "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/repositories"
)

func IncrementCombinationGenerationCounter(c *gin.Context) {
	var bodyRequisition CombinationGenerationCounterEntity.CombinationGenerationCounter
	if err := c.ShouldBindJSON(&bodyRequisition); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": "Error in Json decoding", "statusCode": statusCode})
		return
	}
	numberOfEntries := bodyRequisition.NumberOfEntries
	results, err := CombinationGenerationCounterRepo.FindByNumberOfEntries(numberOfEntries)
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		errorMessage := err.Error()
		c.JSON(statusCode, gin.H{"error": errorMessage, "statusCode": statusCode})
		return
	}
	c.JSON(http.StatusOK, results)
}

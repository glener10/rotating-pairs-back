package CombinationGenerationCounterUseCases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	CombinationGenerationCounterRepo "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/repositories"
)

func IncrementCombinationGenerationCounter(c *gin.Context, bodyRequisition CombinationGenerationCounterEntity.CombinationGenerationCounter) {
	NumberOfInputs := bodyRequisition.NumberOfInputs
	results, err := CombinationGenerationCounterRepo.FindByNumberOfInputs(NumberOfInputs)
	if err != nil {
		results, err = CombinationGenerationCounterRepo.CreateCombinationGenerationCounter(NumberOfInputs)
		if err != nil {
			statusCode := http.StatusUnprocessableEntity
			errorMessage := err.Error()
			c.JSON(statusCode, gin.H{"error": errorMessage, "statusCode": statusCode})
			return
		}
		c.JSON(http.StatusOK, results)
		return
	}
	results, err = CombinationGenerationCounterRepo.IncrementCombinationGenerationCounter(results)
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		errorMessage := err.Error()
		c.JSON(statusCode, gin.H{"error": errorMessage, "statusCode": statusCode})
		return
	}
	c.JSON(http.StatusOK, results)
}

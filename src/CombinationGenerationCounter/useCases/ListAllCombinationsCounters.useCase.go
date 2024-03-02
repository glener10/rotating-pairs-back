package CombinationGenerationCounterUseCases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	CombinationGenerationCounterRepo "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/repositories"
)

func ListAllCombinationsCounters(c *gin.Context) {
	results, err := CombinationGenerationCounterRepo.ListAllCombinationsCounters()
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		errorMessage := err.Error()
		c.JSON(statusCode, gin.H{"error": errorMessage, "statusCode": statusCode})
		return
	}
	c.JSON(http.StatusOK, results)
}

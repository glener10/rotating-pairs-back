package CombinationGenerationCounterController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	IncrementCombinationGenerationCounterUseCase "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/useCases"
)

func IncrementCombinationGenerationCounter(r *gin.Engine) {
	r.POST("/combinationGenerationCounter", func(c *gin.Context) {
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

		IncrementCombinationGenerationCounterUseCase.IncrementCombinationGenerationCounter(c, combination)
	})
}

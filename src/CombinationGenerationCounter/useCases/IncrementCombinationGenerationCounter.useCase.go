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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falha ao decodificar o JSON"})
		return
	}
	numberOfEntries := bodyRequisition.NumberOfEntries
	results, err := CombinationGenerationCounterRepo.FindByNumberOfEntries(numberOfEntries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar os registros"})
		return
	}
	c.JSON(http.StatusOK, results)
}

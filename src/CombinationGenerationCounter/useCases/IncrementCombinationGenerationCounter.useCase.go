package IncrementCombinationGenerationCounterUseCase

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	CombinationGenerationCounterRepo "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/repositories"
)

func IncrementCombinationGenerationCounter(c *gin.Context) {
	repo, err := CombinationGenerationCounterRepo.Connect()

	if err != nil {
		fmt.Println("Erro ao conectar:", err)
		return
	}
	fmt.Println("Conexão bem-sucedida!")
	fmt.Println(repo)
	var bodyRequisition CombinationGenerationCounterEntity.CombinationGenerationCounter
	if err := c.ShouldBindJSON(&bodyRequisition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "fon"})
	}
	c.JSON(200, bodyRequisition)
}

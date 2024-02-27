package CombinationGenerationCounterController

import (
	"github.com/gin-gonic/gin"
	IncrementCombinationGenerationCounterUseCase "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/useCases"
)

func IncrementCombinationGenerationCounter(c *gin.Context) {
	IncrementCombinationGenerationCounterUseCase.IncrementCombinationGenerationCounter(c)
}

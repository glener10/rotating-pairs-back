package CombinationGenerationCounterController

import (
	"github.com/gin-gonic/gin"
	IncrementCombinationGenerationCounterUseCase "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/useCases"
)

func IncrementCombinationGenerationCounter(r *gin.Engine) {
	r.POST("/combinationGenerationCounter", IncrementCombinationGenerationCounterUseCase.IncrementCombinationGenerationCounter)
}

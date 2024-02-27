package routes

import (
	"github.com/gin-gonic/gin"
	CombinationGenerationCounterController "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/controllers"
)

func HandlerRoutes() {
	r := gin.Default()
	CombinationGenerationCounterController.IncrementCombinationGenerationCounter(r)
	r.Run()
}

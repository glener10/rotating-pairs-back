package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	CombinationGenerationCounterController "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/controllers"
	CombinationController "github.com/glener10/rotating-pairs-back/src/Combinations/controllers"
)

func HandlerRoutes() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Use(CorsMiddleware())

	r.POST("/combinationGenerationCounter", CombinationGenerationCounterController.IncrementCombinationGenerationCounter)
	r.GET("/combinationGenerationCounter", CombinationGenerationCounterController.ListAllCombinationsCounters)
	r.POST("/combination", CombinationController.Combination)

	err := r.Run()
	if err != nil {
		fmt.Println("Error to up routes")
		os.Exit(-1)
	}
}

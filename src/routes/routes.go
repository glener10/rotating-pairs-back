package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	CombinationGenerationCounterController "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/controllers"
	CombinationController "github.com/glener10/rotating-pairs-back/src/Combinations/controllers"
	Middlewares "github.com/glener10/rotating-pairs-back/src/routes/middlewares"
)

func HandlerRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	if os.Getenv("ENV") != "development" {
		r.Use(Middlewares.OriginMiddleware())
		r.Use(Middlewares.MethodsMiddleware())
		r.Use(Middlewares.HTTPSOnlyMiddleware())
	}

	r.POST("/combinationGenerationCounter", CombinationGenerationCounterController.IncrementCombinationGenerationCounter)
	r.GET("/combinationGenerationCounter", CombinationGenerationCounterController.ListAllCombinationsCounters)
	r.POST("/combination", CombinationController.Combination)
	return r
}

func Listening(r *gin.Engine) {
	err := r.Run()
	if err != nil {
		fmt.Println("Error to up routes")
		os.Exit(-1)
	}
}

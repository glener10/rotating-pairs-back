package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	CombinationGenerationCounterController "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/controllers"
)

func HandlerRoutes() {
	r := gin.Default()
	CombinationGenerationCounterController.IncrementCombinationGenerationCounter(r)
	err := r.Run()
	if err != nil {
		fmt.Println("Error to up routes")
		os.Exit(-1)
	}
}

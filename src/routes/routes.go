package routes

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	CombinationGenerationCounterController "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/controllers"
	CombinationController "github.com/glener10/rotating-pairs-back/src/Combinations/controllers"
)

func GetAllowedURLs() []string {
	allowedURLsString := os.Getenv("ALLOW_URLS")
	if allowedURLsString == "" {
		return nil
	}
	allowedURLs := strings.Split(allowedURLsString, "|")
	return allowedURLs
}

func HandlerRoutes() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = GetAllowedURLs()
	if config.AllowOrigins == nil {
		fmt.Println("Error to read allowed urls in the Route Handler")
		os.Exit(-1)
	}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	r.POST("/combinationGenerationCounter", CombinationGenerationCounterController.IncrementCombinationGenerationCounter)
	r.GET("/combinationGenerationCounter", CombinationGenerationCounterController.ListAllCombinationsCounters)
	r.POST("/combination", CombinationController.Combination)
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	err := r.Run("3000")
	if err != nil {
		fmt.Println("Error to up routes")
		os.Exit(-1)
	}
}

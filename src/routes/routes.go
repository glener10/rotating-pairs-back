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

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", GetAllowedURLs())
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

func GetAllowedURLs() string {
	allowedURLsString := os.Getenv("ALLOW_URLS")
	if allowedURLsString == "" {
		return "*"
	}
	return allowedURLsString
}

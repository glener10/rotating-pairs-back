package routes

import (
	"fmt"
	"os"
	"strings"

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

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Use(corsMiddleware())

	r.POST("/combinationGenerationCounter", CombinationGenerationCounterController.IncrementCombinationGenerationCounter)
	r.GET("/combinationGenerationCounter", CombinationGenerationCounterController.ListAllCombinationsCounters)
	r.POST("/combination", CombinationController.Combination)

	err := r.Run()
	if err != nil {
		fmt.Println("Error to up routes")
		os.Exit(-1)
	}
}

func corsMiddleware() gin.HandlerFunc {
	allowedUrls := GetAllowedURLs()
	if allowedUrls == nil {
		fmt.Println("Error to read allowed urls in the Route Handler")
		os.Exit(-1)
	}

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://rotatingpairs.online/")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

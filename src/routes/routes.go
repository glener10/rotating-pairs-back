package routes

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	CombinationGenerationCounterController "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/controllers"
	CombinationController "github.com/glener10/rotating-pairs-back/src/Combinations/controllers"
	GenerateController "github.com/glener10/rotating-pairs-back/src/Generate/controllers"
	Middlewares "github.com/glener10/rotating-pairs-back/src/routes/middlewares"
)

func HandlerRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(Middlewares.TimeoutMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     getAllowedURLs(),
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	if os.Getenv("ENV") != "development" {
		rateLimiter := Middlewares.NewRateLimiter(11, time.Minute)
		r.Use(Middlewares.RequestLimitMiddleware(rateLimiter))
		r.Use(Middlewares.AuthMiddleware())
		//r.Use(Middlewares.MethodsMiddleware())
		//r.Use(Middlewares.HTTPSOnlyMiddleware())
	}

	r.POST("/combinationGenerationCounter", CombinationGenerationCounterController.IncrementCombinationGenerationCounter)
	r.GET("/combinationGenerationCounter", CombinationGenerationCounterController.ListAllCombinationsCounters)
	r.POST("/combination", CombinationController.Combination)
	r.POST("/generate", GenerateController.Generate)
	return r
}

func Listening(r *gin.Engine) {
	err := r.Run()
	if err != nil {
		fmt.Println("Error to up routes")
		os.Exit(-1)
	}
}

func getAllowedURLs() []string {
	if os.Getenv("ENV") == "development" {
		return []string{"*"}
	}
	allowedURLsString := os.Getenv("ALLOW_URLS")
	if allowedURLsString == "" {
		return nil
	}
	allowedUrlsString := strings.Split(allowedURLsString, "|")
	return allowedUrlsString
}

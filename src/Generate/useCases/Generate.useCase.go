package GenerateUseCases

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	CombinationRepo "github.com/glener10/rotating-pairs-back/src/Combinations/repositories"
	GenerateRequestDto "github.com/glener10/rotating-pairs-back/src/Generate/dtos"
)

func Generate(c *gin.Context, combinationRequest GenerateRequestDto.GenerateRequest) {
	NumberOfInputs := combinationRequest.NumberOfInputs
	_, err := CombinationRepo.FindCombination(NumberOfInputs)
	if err == nil {
		statusCode := http.StatusOK
		c.JSON(statusCode, gin.H{"error": "There is already a transaction with this number of entries", "statusCode": statusCode})
		return
	}
	statusCode := http.StatusOK
	c.JSON(statusCode, gin.H{"error": "Trying generating combinations without repetitions...", "statusCode": statusCode})
	if shouldReturn() {
		fmt.Println("Error to generate Combination")
		return
	}
	fmt.Println("Combination Generated")
}

func shouldReturn() bool {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomNumber := rng.Intn(101)
	return randomNumber < 50
}

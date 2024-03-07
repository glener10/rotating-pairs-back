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
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(101) // Gera um número aleatório entre 0 e 100
	if randomNumber < 50 {
		return true
	}
	return false
}

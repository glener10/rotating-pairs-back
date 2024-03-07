package GenerateController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	GenerateRequestDto "github.com/glener10/rotating-pairs-back/src/Generate/dtos"
	GenerateUseCases "github.com/glener10/rotating-pairs-back/src/Generate/useCases"
)

func Generate(c *gin.Context) {
	var generateRequest GenerateRequestDto.GenerateRequest
	if err := c.ShouldBindJSON(&generateRequest); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "statusCode": statusCode})
		return
	}
	if err := GenerateRequestDto.Validate(&generateRequest); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	GenerateUseCases.Generate(c, generateRequest)
}

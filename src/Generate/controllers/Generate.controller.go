package GenerateController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	GenerateUseCases "github.com/glener10/rotating-pairs-back/src/Generate/useCases"
	CombinationRequestDto "github.com/glener10/rotating-pairs-back/src/common/interfaces"
)

func Generate(c *gin.Context) {
	var combinationRequest CombinationRequestDto.CombinationRequest
	if err := c.ShouldBindJSON(&combinationRequest); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "statusCode": statusCode})
		return
	}

	GenerateUseCases.Generate(c, combinationRequest)
}

package Middlewares

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	CombinationController "github.com/glener10/rotating-pairs-back/src/Combinations/controllers"
	"github.com/stretchr/testify/assert"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	return r
}

func TestMethodNotAllowedDellete(t *testing.T) {
	r := SetupRoutes()
	r.Use(MethodsMiddleware())
	r.DELETE("/combination", CombinationController.Combination)
	req, _ := http.NewRequest("DELETE", "/combination", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := ErrorResponse{
		Error:      "Method not allowed",
		StatusCode: 403,
	}

	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, expected, actual, "Should return 'Method not allowed' and 403 if the requisition is a DELETE method")
}

func TestMethodNotAllowedPatch(t *testing.T) {
	r := SetupRoutes()
	r.Use(MethodsMiddleware())
	r.PATCH("/combination", CombinationController.Combination)
	req, _ := http.NewRequest("PATCH", "/combination", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := ErrorResponse{
		Error:      "Method not allowed",
		StatusCode: 403,
	}

	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, expected, actual, "Should return 'Method not allowed' and 403 if the requisition is a PATCH method")
}

func TestMethodOptionsReturn(t *testing.T) {
	r := SetupRoutes()
	r.Use(MethodsMiddleware())
	r.OPTIONS("/combination", CombinationController.Combination)
	req, _ := http.NewRequest("OPTIONS", "/combination", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Code, http.StatusOK, "Should return a 200 code")
}

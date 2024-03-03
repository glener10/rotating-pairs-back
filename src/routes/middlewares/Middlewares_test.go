package Middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	CombinationGenerationCounterController "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/controllers"
	CombinationController "github.com/glener10/rotating-pairs-back/src/Combinations/controllers"
	Utils "github.com/glener10/rotating-pairs-back/src/common/utils"
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

func TestOnlyHttps(t *testing.T) {
	r := SetupRoutes()
	r.Use(HTTPSOnlyMiddleware())
	r.GET("/combinationGenerationCounter", CombinationController.Combination)
	req, _ := http.NewRequest("GET", "/combinationGenerationCounter", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := ErrorResponse{
		Error:      "HTTPS only, your protocol is: HTTP/1.1",
		StatusCode: 403,
	}

	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, actual, expected, "Should return 'HTTPS only' and 403 if the requisition its not HTTPS")
}

func TestRateLimiter(t *testing.T) {
	r := SetupRoutes()
	rateLimiter := NewRateLimiter(1, time.Minute)
	r.Use(RequestLimitMiddleware(rateLimiter))
	r.GET("/combinationGenerationCounter", CombinationController.Combination)
	req, _ := http.NewRequest("GET", "/combinationGenerationCounter", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	r.ServeHTTP(response, req)

	expected := ErrorResponse{
		Error:      "Too Many Requests",
		StatusCode: 429,
	}

	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, actual, expected, "Should return 'Too Many Requests' and 429 if the requisition pass the rate limiter")
}

func TestAuthWithNoToken(t *testing.T) {
	r := SetupRoutes()
	r.Use(AuthMiddleware())
	r.GET("/combinationGenerationCounter", CombinationController.Combination)
	req, _ := http.NewRequest("GET", "/combinationGenerationCounter", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := ErrorResponse{
		Error:      "Token not Provided",
		StatusCode: 422,
	}

	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, actual, expected, "Should return a 422 because de token is not informed")
}

func TestAuthWithInvalidToken(t *testing.T) {
	r := SetupRoutes()
	r.Use(AuthMiddleware())
	r.GET("/combinationGenerationCounter", CombinationController.Combination)
	req, _ := http.NewRequest("GET", "/combinationGenerationCounter", nil)
	req.Header.Set("Authorization", "Bearer TOKEN_INVALIDO")
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := ErrorResponse{
		Error:      "Invalid Token",
		StatusCode: 401,
	}

	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, actual, expected, "Should return a 422 because de token is not informed")
}

func TestAuthWithValidToken(t *testing.T) {
	if err := Utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("Error to load environment variables: %s", err.Error())
	}
	r := SetupRoutes()
	r.Use(AuthMiddleware())
	r.GET("/combinationGenerationCounter", CombinationGenerationCounterController.ListAllCombinationsCounters)
	req, _ := http.NewRequest("GET", "/combinationGenerationCounter", nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SECRET"))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Code, 200, "Should return code 200 with a valid token")
}

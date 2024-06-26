package Middlewares

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

	var errorResponseDecoded ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&errorResponseDecoded)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, errorResponseDecoded, expected, "Should return 'Too Many Requests' and 429 if the requisition pass the rate limiter")
}

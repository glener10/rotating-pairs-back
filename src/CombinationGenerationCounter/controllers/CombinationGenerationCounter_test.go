package CombinationGenerationCounterController

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	CombinationGenerationCounterEntity "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/entities"
	CombinationGenerationCounterUtils "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/utils"
	Utils "github.com/glener10/rotating-pairs-back/src/common/utils"
	"github.com/stretchr/testify/assert"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

type BodyRequest struct {
	NumberOfEntries int16 `json:"NumberOfEntries"`
}

func TestMain(m *testing.M) {
	if err := Utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("Error to load environment variables: %s", err.Error())
	}
	if err := CombinationGenerationCounterUtils.Truncate(); err != nil {
		log.Fatalf("Error to exec truncate method before repository tests execution: %s", err.Error())
	}
	exitCode := m.Run()
	if err := CombinationGenerationCounterUtils.CleanCollection(); err != nil {
		log.Fatalf("Error to exec cleaning collection after repository tests execution: %s", err.Error())
	}
	os.Exit(exitCode)
}

func SetupRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestRouteWithoutBody(t *testing.T) {
	r := SetupRoutes()
	r.POST("/combinationGenerationCounter", IncrementCombinationGenerationCounter)
	req, _ := http.NewRequest("POST", "/combinationGenerationCounter", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := ErrorResponse{
		Error:      "Invalid request body",
		StatusCode: 422,
	}

	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, expected, actual, "Should return 'Invalid request body' and 422 if the requisition doenst have a body")
}

func TestRouteWitNumberOfInputsMoreThanTwenty(t *testing.T) {
	r := SetupRoutes()
	r.POST("/combinationGenerationCounter", IncrementCombinationGenerationCounter)

	body := BodyRequest{
		NumberOfEntries: 25,
	}
	bodyConverted, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/combinationGenerationCounter", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := ErrorResponse{
		Error:      "NumberOfEntries is more than 20 or less than 1",
		StatusCode: 422,
	}

	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, expected, actual, "Should return 'NumberOfEntries is more than 20 or less than 1' and 422 if the NumberOfInputs in body is more than 20")
}

func TestRouteWitNumberOfInputsLessThanZero(t *testing.T) {
	r := SetupRoutes()
	r.POST("/combinationGenerationCounter", IncrementCombinationGenerationCounter)

	body := BodyRequest{
		NumberOfEntries: -1,
	}
	bodyConverted, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/combinationGenerationCounter", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := ErrorResponse{
		Error:      "NumberOfEntries is more than 20 or less than 1",
		StatusCode: 422,
	}

	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, expected, actual, "Should return 'NumberOfEntries is more than 20 or less than 1' and 422 if the NumberOfInputs in body is less tha 0")
}

func TestRouteSuccessCase(t *testing.T) {
	r := SetupRoutes()
	r.POST("/combinationGenerationCounter", IncrementCombinationGenerationCounter)

	body := BodyRequest{
		NumberOfEntries: 2,
	}
	bodyConverted, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/combinationGenerationCounter", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	two := int32(2)
	expected := CombinationGenerationCounterEntity.CombinationGenerationCounter{
		Count:           &two,
		NumberOfEntries: 2,
	}

	var actual CombinationGenerationCounterEntity.CombinationGenerationCounter
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, expected, actual, "Should return expected object in the success case")
}

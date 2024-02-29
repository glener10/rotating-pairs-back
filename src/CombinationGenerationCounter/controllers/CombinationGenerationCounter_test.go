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
	CombinationGenerationCounterRepo "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/repositories"
	CombinationGenerationCounterUtils "github.com/glener10/rotating-pairs-back/src/CombinationGenerationCounter/utils"
	Utils "github.com/glener10/rotating-pairs-back/src/common/utils"
	"github.com/stretchr/testify/assert"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

type BodyRequest struct {
	NumberOfInputs int16 `json:"NumberOfInputs"`
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

func TestIncrementRouteWithoutBody(t *testing.T) {
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

func TestIncrementRouteWitNumberOfInputsMoreThanTwenty(t *testing.T) {
	r := SetupRoutes()
	r.POST("/combinationGenerationCounter", IncrementCombinationGenerationCounter)

	body := BodyRequest{
		NumberOfInputs: 25,
	}
	bodyConverted, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/combinationGenerationCounter", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := ErrorResponse{
		Error:      "NumberOfInputs is more than 20 or less than 2",
		StatusCode: 422,
	}

	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, expected, actual, "Should return 'NumberOfInputs is more than 20 or less than 2' and 422 if the NumberOfInputs in body is more than 20")
}

func TestIncrementRouteWitNumberOfInputsLessThanZero(t *testing.T) {
	r := SetupRoutes()
	r.POST("/combinationGenerationCounter", IncrementCombinationGenerationCounter)

	body := BodyRequest{
		NumberOfInputs: -1,
	}
	bodyConverted, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/combinationGenerationCounter", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := ErrorResponse{
		Error:      "NumberOfInputs is more than 20 or less than 2",
		StatusCode: 422,
	}

	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, expected, actual, "Should return 'NumberOfInputs is more than 20 or less than 2' and 422 if the NumberOfInputs in body is less tha 0")
}

func TestIncrementRouteSuccessCase(t *testing.T) {
	r := SetupRoutes()
	r.POST("/combinationGenerationCounter", IncrementCombinationGenerationCounter)

	body := BodyRequest{
		NumberOfInputs: 2,
	}
	bodyConverted, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/combinationGenerationCounter", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := CombinationGenerationCounterEntity.CombinationGenerationCounter{
		Count:          2,
		NumberOfInputs: 2,
	}

	var actual CombinationGenerationCounterEntity.CombinationGenerationCounter
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, expected, actual, "Should return expected object in the success case")
}

func TestListAllCombinationsCountersRouteSuccessCase(t *testing.T) {
	if err := CombinationGenerationCounterUtils.CleanCollection(); err != nil {
		log.Fatalf("Error to exec cleaning collection after repository tests execution: %s", err.Error())
	}

	_, _ = CombinationGenerationCounterRepo.CreateCombinationGenerationCounter(1)
	_, _ = CombinationGenerationCounterRepo.CreateCombinationGenerationCounter(2)
	_, _ = CombinationGenerationCounterRepo.IncrementCombinationGenerationCounter(2)
	_, _ = CombinationGenerationCounterRepo.IncrementCombinationGenerationCounter(2)

	r := SetupRoutes()
	r.GET("/combinationGenerationCounter", ListAllCombinationsCounters)
	req, _ := http.NewRequest("GET", "/combinationGenerationCounter", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	firstExpectededObject := CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfInputs: 2,
		Count:          3,
	}
	secondExpectededObject := CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfInputs: 1,
		Count:          1,
	}

	var actual []CombinationGenerationCounterEntity.CombinationGenerationCounter
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, len(actual), 2, "Slice length should be equal")
	assert.Equal(t, (actual)[0], firstExpectededObject, "Element should be equal")
	assert.Equal(t, (actual)[1], secondExpectededObject, "Element should be equal")
}

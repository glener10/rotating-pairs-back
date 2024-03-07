package CombinationController

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
	CombinationEntity "github.com/glener10/rotating-pairs-back/src/Combinations/entities"
	CombinationRepo "github.com/glener10/rotating-pairs-back/src/Combinations/repositories"
	CombinationRequestDto "github.com/glener10/rotating-pairs-back/src/common/dtos"
	Utils "github.com/glener10/rotating-pairs-back/src/common/utils"
	"github.com/stretchr/testify/assert"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func TestMain(m *testing.M) {
	if err := Utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("Error to load environment variables: %s", err.Error())
	}
	if err := CombinationRepo.Truncate(); err != nil {
		log.Fatalf("Error to exec truncate method before controller Combinations tests execution: %s", err.Error())
	}
	exitCode := m.Run()
	if err := CombinationRepo.CleanCollection(); err != nil {
		log.Fatalf("Error to exec cleaning collection after controller Combinations tests execution: %s", err.Error())
	}
	os.Exit(exitCode)
}

func SetupRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestCombinationRouteWithoutBody(t *testing.T) {
	r := SetupRoutes()
	r.POST("/combination", Combination)
	req, _ := http.NewRequest("POST", "/combination", nil)
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

func TestCombinationRouteWitNumberOfInputsMoreThanTwenty(t *testing.T) {
	r := SetupRoutes()
	r.POST("/combination", Combination)
	body := CombinationRequestDto.CombinationRequest{
		NumberOfInputs: 25,
	}
	bodyConverted, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/combination", bytes.NewBuffer(bodyConverted))
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

func TestCombinationRouteWitNumberOfInputsLessThanTwo(t *testing.T) {
	r := SetupRoutes()
	r.POST("/combination", Combination)
	body := CombinationRequestDto.CombinationRequest{
		NumberOfInputs: -1,
	}
	bodyConverted, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/combination", bytes.NewBuffer(bodyConverted))
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

func TestCombinationRouteSuccessCase(t *testing.T) {
	r := SetupRoutes()
	r.POST("/combination", Combination)
	body := CombinationRequestDto.CombinationRequest{
		NumberOfInputs: 2,
	}
	bodyConverted, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/combination", bytes.NewBuffer(bodyConverted))

	_, _ = CombinationGenerationCounterRepo.CreateCombinationGenerationCounter(2)

	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	combinationCounterAfterUseCase, _ := CombinationGenerationCounterRepo.FindByNumberOfInputs(2)

	expectedCombinationCounter := &CombinationGenerationCounterEntity.CombinationGenerationCounter{
		NumberOfInputs: 2,
		Count:          2,
	}

	expectedObject := &CombinationEntity.Combination{
		NumberOfInputs:                2,
		NumberOfSprints:               1,
		NumberOfCombinationsPerSprint: 1,
		Sprints: []CombinationEntity.Sprint{
			{
				Combinations: []CombinationEntity.Pair{
					{
						PairOne: 0,
						PairTwo: 1,
					},
				},
			},
		},
	}

	var actual *CombinationEntity.Combination
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, combinationCounterAfterUseCase, expectedCombinationCounter, "Should increment the Combination Counter after useCase executation")
	assert.Equal(t, actual, expectedObject, "Should return the expected object")
}

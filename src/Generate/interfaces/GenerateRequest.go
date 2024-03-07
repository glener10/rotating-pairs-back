package GenerateRequestDto

import (
	"errors"
)

type GenerateRequest struct {
	NumberOfInputs                                                                         int16 `json:"NumberOfInputs"`
	MAX_NUMBER_OF_SHUFFLING_ARRAY_OF_POSSIBLE_COMBINATIONS_BEFORE_ALTERNING_SHUFFLE_METHOD int
	MAX_NUMBER_OF_ALTERNING_SHUFFLE_METHOD                                                 int
}

func Validate(combination *GenerateRequest) error {
	if combination.NumberOfInputs < 2 {
		return errors.New("NumberOfInputs is less than 2")
	}
	if combination.MAX_NUMBER_OF_SHUFFLING_ARRAY_OF_POSSIBLE_COMBINATIONS_BEFORE_ALTERNING_SHUFFLE_METHOD == 0 {
		combination.MAX_NUMBER_OF_SHUFFLING_ARRAY_OF_POSSIBLE_COMBINATIONS_BEFORE_ALTERNING_SHUFFLE_METHOD = 50000
	}
	if combination.MAX_NUMBER_OF_ALTERNING_SHUFFLE_METHOD == 0 {
		combination.MAX_NUMBER_OF_ALTERNING_SHUFFLE_METHOD = 10000
	}
	return nil
}

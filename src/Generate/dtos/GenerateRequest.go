package GenerateRequestDto

import (
	"errors"
)

type GenerateRequest struct {
	NumberOfInputs int16 `json:"NumberOfInputs"`
}

func Validate(combination *GenerateRequest) error {
	if combination.NumberOfInputs < 2 {
		return errors.New("NumberOfInputs is less than 2")
	}
	return nil
}

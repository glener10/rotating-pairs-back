package CombinationRequestDto

import "errors"

type CombinationRequest struct {
	NumberOfInputs int16 `json:"NumberOfInputs"`
}

func Validate(combination *CombinationRequest) error {
	if combination.NumberOfInputs < 2 || combination.NumberOfInputs > 20 {
		return errors.New("NumberOfInputs is more than 20 or less than 2")
	}
	return nil
}

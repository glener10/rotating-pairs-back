package CombinationGenerationCounterEntity

import "errors"

type CombinationGenerationCounter struct {
	NumberOfInputs int16  `json:"NumberOfInputs"`
	Count          *int32 `json:"Count"`
}

func Validate(combination *CombinationGenerationCounter) error {
	if combination.NumberOfInputs < 1 || combination.NumberOfInputs > 20 {
		return errors.New("NumberOfInputs is more than 20 or less than 1")
	}
	return nil
}

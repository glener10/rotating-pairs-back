package CombinationGenerationCounterEntity

import "errors"

type CombinationGenerationCounter struct {
	NumberOfEntries int16  `json:"NumberOfEntries"`
	Count           *int32 `json:"Count"`
}

func Validate(combination *CombinationGenerationCounter) error {
	if combination.NumberOfEntries < 1 || combination.NumberOfEntries > 20 {
		return errors.New("NumberOfEntries is more than 20 or less than 1")
	}
	return nil
}

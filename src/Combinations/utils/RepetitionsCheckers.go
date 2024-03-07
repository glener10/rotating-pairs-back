package CombinationUtils

import (
	"fmt"

	CombinationEntity "github.com/glener10/rotating-pairs-back/src/Combinations/entities"
)

/* func DoAllChecksInSprint(sprints []CombinationEntity.Sprint, numberOfInputs int) bool {
	if !checkIfAllIndexesAreValid(sprints, numberOfInputs) {
		return false
	}
	if !checkIfThereIsARepeatedCombination(sprints) {
		return false
	}
	if !checkIfAnyInputFromThePairIsRepeatedInTheCombinationsOfASprint(sprints) {
		return false
	}
	indexArrayWithNumberInputs := generateIndexArrayWithSizeOfNewEntry(numberOfInputs)
	numberOfInputsIsOdd := checkIfArrayIsOdd(indexArrayWithNumberInputs)
	numberOfCombinationPerSprint := returnNumberOfCombinationPerSprintRoundedDown(numberOfInputsIsOdd, indexArrayWithNumberInputs)
	if !checkIfAllSprintsHaveAValidNumberOfCombinations(sprints, numberOfCombinationPerSprint) {
		return false
	}
	numberOfSprints := returnNumberOfSprints(numberOfInputsIsOdd, indexArrayWithNumberInputs)
	if !checkIfAllCombinationsHaveAValidNumberOfSprint(sprints, numberOfSprints) {
		return false
	}
	return true
} */

func CheckIfThereIsARepeatedCombination(Sprints []CombinationEntity.Sprint) bool {
	for indexSprint := 0; indexSprint < len(Sprints); indexSprint++ {
		sprint := Sprints[indexSprint]
		for indexCombinations := 0; indexCombinations < len(sprint.Combinations); indexCombinations++ {
			combination := sprint.Combinations[indexCombinations]
			if !CheckOneCombination(combination, indexSprint, indexCombinations, Sprints) {
				return false
			}
		}
	}
	return true
}

func CheckOneCombination(combination CombinationEntity.Pair, indexSprintYourself, indexCombinationsYourself int, sprints []CombinationEntity.Sprint) bool {
	for indexSprint := 0; indexSprint < len(sprints); indexSprint++ {
		sprint := sprints[indexSprint]

		for indexCombinations := 0; indexCombinations < len(sprint.Combinations); indexCombinations++ {
			combinationForTesting := sprint.Combinations[indexCombinations]

			if indexSprint != indexSprintYourself || indexCombinations != indexCombinationsYourself {
				if combination.PairOne == combination.PairTwo || combinationForTesting.PairOne == combinationForTesting.PairTwo {
					if combination.PairOne == combinationForTesting.PairOne && combination.PairOne == combinationForTesting.PairTwo && combination.PairTwo == combinationForTesting.PairOne && combination.PairTwo == combinationForTesting.PairTwo {
						fmt.Printf("Sprint[%d]: %d - %d   With  Sprint[%d]: %d - %d\n", indexSprintYourself, combination.PairOne, combination.PairTwo, indexSprint, combinationForTesting.PairOne, combinationForTesting.PairTwo)
						return false
					}
				} else {
					if (combination.PairOne == combinationForTesting.PairOne || combination.PairOne == combinationForTesting.PairTwo) && (combination.PairTwo == combinationForTesting.PairOne || combination.PairTwo == combinationForTesting.PairTwo) {
						fmt.Printf("Sprint[%d]: %d - %d   With  Sprint[%d]: %d - %d\n", indexSprintYourself, combination.PairOne, combination.PairTwo, indexSprint, combinationForTesting.PairOne, combinationForTesting.PairTwo)
						return false
					}
				}
			}
		}
	}
	return true
}

func CheckIfAllIndexesAreValid(sprints []CombinationEntity.Sprint, numberOfInputs int) bool {
	for indexSprint := 0; indexSprint < len(sprints); indexSprint++ {
		sprint := sprints[indexSprint]
		for indexCombinations := 0; indexCombinations < len(sprint.Combinations); indexCombinations++ {
			combination := sprint.Combinations[indexCombinations]
			if combination.PairOne >= numberOfInputs || combination.PairTwo >= numberOfInputs {
				return false
			}
		}
	}
	return true
}

func CheckIfAnyInputFromThePairIsRepeatedInTheCombinationsOfASprint(sprints []CombinationEntity.Sprint) bool {
	for indexSprint := 0; indexSprint < len(sprints); indexSprint++ {
		sprint := sprints[indexSprint]
		for indexCombinations := 0; indexCombinations < len(sprint.Combinations); indexCombinations++ {
			combination := sprint.Combinations[indexCombinations]
			if !CheckAnyPairDontRepeatInTheCombinations(combination, indexCombinations, sprint.Combinations) {
				return false
			}
		}
	}
	return true
}

func CheckAnyPairDontRepeatInTheCombinations(combination CombinationEntity.Pair, indexCombinationsYourself int, combinations []CombinationEntity.Pair) bool {
	for indexCombinations := 0; indexCombinations < len(combinations); indexCombinations++ {
		combinationForTesting := combinations[indexCombinations]
		if indexCombinations != indexCombinationsYourself {
			if combination.PairOne == combinationForTesting.PairOne ||
				combination.PairOne == combinationForTesting.PairTwo ||
				combination.PairTwo == combinationForTesting.PairOne ||
				combination.PairTwo == combinationForTesting.PairTwo {
				return false
			}
		}
	}
	return true
}

func CheckIfAllSprintsHaveAValidNumberOfCombinations(sprints []CombinationEntity.Sprint, numberOfCombinationPerSprint int) bool {
	for indexSprint := 0; indexSprint < len(sprints); indexSprint++ {
		if len(sprints[indexSprint].Combinations) != numberOfCombinationPerSprint {
			return false
		}
	}
	return true
}

func CheckIfAllCombinationsHaveAValidNumberOfSprint(sprints []CombinationEntity.Sprint, numberOfSprints int) bool {
	return len(sprints) == numberOfSprints
}

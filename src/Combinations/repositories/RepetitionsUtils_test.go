package CombinationRepo

import (
	"fmt"

	CombinationEntity "github.com/glener10/rotating-pairs-back/src/Combinations/entities"
)

func returnAllCombinations() []*CombinationEntity.Combination {
	var combinations []*CombinationEntity.Combination

	for i := 2; i <= 3; i++ {
		found, err := FindCombination(int16(i))
		if err != nil {
			continue
		}
		combinations = append(combinations, found)
	}
	return combinations
}

func checkIfThereIsARepeatedCombination(Sprints []CombinationEntity.Sprint) bool {
	for indexSprint := 0; indexSprint < len(Sprints); indexSprint++ {
		sprint := Sprints[indexSprint]
		for indexCombinations := 0; indexCombinations < len(sprint.Combinations); indexCombinations++ {
			combination := sprint.Combinations[indexCombinations]
			if !checkOneCombination(combination, indexSprint, indexCombinations, Sprints) {
				return false
			}
		}
	}
	return true
}

func checkOneCombination(combination CombinationEntity.Pair, indexSprintYourself, indexCombinationsYourself int, sprints []CombinationEntity.Sprint) bool {
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

func checkIfAllIndexesAreValid(sprints []CombinationEntity.Sprint, numberOfInputs int) bool {
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

func checkIfAnyInputFromThePairIsRepeatedInTheCombinationsOfASprint(sprints []CombinationEntity.Sprint) bool {
	for indexSprint := 0; indexSprint < len(sprints); indexSprint++ {
		sprint := sprints[indexSprint]
		for indexCombinations := 0; indexCombinations < len(sprint.Combinations); indexCombinations++ {
			combination := sprint.Combinations[indexCombinations]
			if !checkAnyPairDontRepeatInTheCombinations(combination, indexCombinations, sprint.Combinations) {
				return false
			}
		}
	}
	return true
}

func checkAnyPairDontRepeatInTheCombinations(combination CombinationEntity.Pair, indexCombinationsYourself int, combinations []CombinationEntity.Pair) bool {
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

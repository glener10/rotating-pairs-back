package CombinationEntity

type Combination struct {
	NumberOfInputs                int      `json:"NumberOfInputs"`
	NumberOfSprints               int      `json:"NumberOfSprints"`
	NumberOfCombinationsPerSprint int      `json:"NumberOfCombinationsPerSprint"`
	Sprints                       []Sprint `json:"Sprints"`
}

type Sprint struct {
	Combinations []Pair `json:"Combinations"`
}

type Pair struct {
	PairOne int `json:"PairOne"`
	PairTwo int `json:"PairTwo"`
}

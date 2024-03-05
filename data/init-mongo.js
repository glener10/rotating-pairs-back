db.createCollection('TotalCombinationGenerationAccordingNumberOfInputs');

db.TotalCombinationGenerationAccordingNumberOfInputs.insertMany([
  { NumberOfInputs: 2, Count: 0 },
  { NumberOfInputs: 3, Count: 0 }
]);

db.createCollection('Combinations');

db.Combinations.insertMany([
  {
    NumberOfInputs: 2,
    NumberOfSprints: 1,
    NumberOfCombinationsPerSprint: 1,
    Sprints: [
      {
        Combinations: [
          {
            PairOne: 0,
            PairTwo: 1
          }
        ]
      }
    ]
  },
  {
    NumberOfInputs: 3,
    NumberOfSprints: 3,
    NumberOfCombinationsPerSprint: 2,
    Sprints: [
      {
        Combinations: [
          {
            PairOne: 0,
            PairTwo: 1
          },
          {
            PairOne: 2,
            PairTwo: 2
          }
        ]
      },
      {
        Combinations: [
          {
            PairOne: 0,
            PairTwo: 2
          },
          {
            PairOne: 1,
            PairTwo: 1
          }
        ]
      },
      {
        Combinations: [
          {
            PairOne: 1,
            PairTwo: 2
          },
          {
            PairOne: 0,
            PairTwo: 0
          }
        ]
      }
    ]
  }
]);
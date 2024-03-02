SCENARIO: Get Combinations according number of inputs
GIVEN A request was sent containing the number of inputs
WHEN User clicks on the "Generate Random Combinations" button in the Front-End
THEN It will return the combinations according the specified number of inputs and exec 'IncrementCombinationGeneration' useCase

SCENARIO: Dont Return Combinations
GIVEN A request was sent without the number of inputs
WHEN User clicks on the "Generate Random Combinations" button in the Front-End without number of inputs or the number of inputs is less than 2
THEN It will return a 422 and a error message

SCENARIO: Dont Return Combinations
GIVEN A request was sent without the number of inputs
WHEN User clicks on the "Generate Random Combinations" button in the Front-End and the number of inputs is more than 20
THEN It will return a 422 and a error message

OBRIGATIONS:

- Implementing testing for both scenarios
- Create Swagger Documentation

SCENARIO: Create a Generation of Combination
GIVEN A request was sent containing the number of entries in the body and a date
WHEN User clicks on the "Generate Random Combinations" button in the Front-End
THEN It will create a GeneratedCombination with number of inputs and a date in 'GeneratedCombinations' collection

SCENARIO: Dont Create a Generation of Combination
GIVEN A request was sent without the number of entries or a date in the body
WHEN User clicks on the "Generate Random Combinations" button in the Front-End
THEN Will return a 422 and a error message

SCENARIO: Dont Create a Generation of Combination
GIVEN A request was sent and the number of entries in the body is more than 20
WHEN User clicks on the "Generate Random Combinations" button in the Front-End
THEN Will return a 422 and a error message

OBRIGATIONS:

- Implementing testing for both scenarios
- Create Swagger Documentation

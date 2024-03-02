SCENARIO: Increment the number of total of combination generation according to the number of inputs
GIVEN A request was sent containing the number of inputs in the body
WHEN User clicks on the "Generate Random Combinations" button in the Front-End
THEN It will increase the total value of combinations that were generated with that same number of inputs in the 'TotalCombinationGenerationAccordingNumberOfInputs' collection

SCENARIO: Do not increment the number of total of combination generation because the body not contains the number of inputs or contains and its less than 2
GIVEN A request was sent without the number of inputs in the body or the number of inputs is less than 2
WHEN User clicks on the "Generate Random Combinations" button in the Front-End without inputs or the inputs is less than 2
THEN Will return a 422 and a message

SCENARIO: Do not increment the number of total of combination generation because the number of inputs is more than 20
GIVEN A request was sent and the number of inputs is more than 20
WHEN User clicks on the "Generate Random Combinations" button in the Front-End
THEN Will return a 422 and a message

OBRIGATIONS:

- Implementing testing for both scenarios
- Create Swagger Documentation

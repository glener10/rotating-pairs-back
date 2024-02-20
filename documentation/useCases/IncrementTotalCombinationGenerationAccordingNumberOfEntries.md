SCENARIO: Increment the number of total of combination generation according to the number of entries
GIVEN A request was sent containing the number of entries in the body
WHEN User clicks on the "Generate Random Combinations" button in the Front-End
THEN It will increase the total value of combinations that were generated with that same number of entries in the 'TotalCombinationGenerationAccordingNumberOfEntries' collection

SCENARIO: Do not increment the number of total of combination generation because the body not contains the number of entries or contains and its less than 2
GIVEN A request was sent without the number of entries in the body or the number of entries is less than 2
WHEN User clicks on the "Generate Random Combinations" button in the Front-End without entries or the entries is less than 2
THEN Will return a 422 and a message

OBRIGATIONS:

- Implementing testing for both scenarios
- Create Swagger Documentation
- Create flow with Draw.io

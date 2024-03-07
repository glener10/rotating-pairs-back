SCENARIO: Valid combination with input number X already exists
GIVEN A POST request was sent containing the number of inputs
WHEN Request made directly to the Back-End (Not open to the front)
THEN It will return the message "There is already a transaction with this number of entries" and an OK code

SCENARIO: NumberOfInputs less than 2
GIVEN A POST request was sent containing the number of inputs < 2
WHEN Request made directly to the Back-End (Not open to the front)
THEN It will return the message "NumberOfInputs is less than 2" and an 422 code

SCENARIO: Successfully trying to generate combination that does not exist
GIVEN A POST request was sent containing the number of inputs
WHEN Request made directly to the Back-End (Not open to the front)
THEN It will return that it is trying to generate combinations with the number of entries sent, and will generate the combination and run the tests that check repetitions with success, saving the new combination in the database and sending a success email

SCENARIO: Unsuccessfully trying to generate combination that does not exist
GIVEN A POST request was sent containing the number of inputs
WHEN Request made directly to the Back-End (Not open to the front)
THEN It will return that it is trying to generate combinations with the number of entries sent, there will be a processing error or a repeating combination, and then it will send a warning email

OBRIGATIONS:

- Implementing testing
- Create Swagger Documentation

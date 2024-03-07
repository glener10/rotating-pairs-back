# [v1.1.0] - XX/XX/XXXX (dd/mm/YYYY)

- New route (POST /generate) that tries to generate combination with input number X without repetitions
- New test switches to check for repetitions in combinations and invalid numbers

# [v1.0.4] - 05/03/2024 (dd/mm/YYYY)

- Fix allowed headers to front production application
- POST /combination refactoring to create a 'CombinationGenerationCounter' if it doesn't exist
- Updating _init-mongo.js_ with basic values

# [v1.0.3] - 03/03/2024 (dd/mm/YYYY)

- Configuring Methods Allowed
- Create middleware to only accepts HTTPS protocols
- Create Rater Limit middleware
- Implement Auth Middleware

# [v1.0.2] - 02/03/2024 (dd/mm/YYYY)

- Configuring CORS policy

# [v1.0.1] - 02/03/2024 (dd/mm/YYYY)

- Changing location of the _Dockerfile_ and _docker-compose_ MongoDB files, because the conflicting in _RailWay_ application
- Updating Documentation
- Moving _main.go_ file to root folder

# [v1.0.0] - 02/03/2024 (dd/mm/YYYY)

- Config setup tests
- Implementation of 'IncrementCombinationGenerationCounter' useCase and integration and E2E tests
- Create CI/CD GitHub Actions procedure
- Implementation of 'ListAllCombinationsCounters' useCase and integration and E2E tests
- Implementation of 'Combination' useCase and integration and E2E tests
- CORS configuration

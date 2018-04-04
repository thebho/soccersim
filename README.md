[![CircleCI](https://circleci.com/gh/thebho/soccersim.svg?style=svg)](https://circleci.com/gh/thebho/soccersim)
[![Go Report Card](https://goreportcard.com/badge/github.com/thebho/soccersim)](https://goreportcard.com/report/github.com/thebho/soccersim)

# Soccer Sim
### Run
1. Add a .env file with `APP_DB_USER`, `APP_DB_PASS`, `APP_DB_NAME`, and `APP_DB_ADDR` for a postgres database
1. Add Goose binary from https://github.com/pressly/goose
1. Run `goose postgres $APP_DB_ADDR up` in ROOT/data directory
1. Start server with `go run main.go`


#### The main functions of this project are:
1. Create and store soccer teams and their seasons
2. Create new seasons and generate a randomized 38 match schedule for each team
3. Allow for simulating the results of a season (currently based on a random selected goal total for the home and away team)
4. Save the simulated matches and team record/goals scored
5. Expose this functionality in a REST API for SoccerSimUI and testing
6. Keep full test coverage and with 12 factor design

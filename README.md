[![CircleCI](https://circleci.com/gh/thebho/soccersim.svg?style=svg)](https://circleci.com/gh/thebho/soccersim)
[![Go Report Card](https://goreportcard.com/badge/github.com/thebho/soccersim)](https://goreportcard.com/report/github.com/thebho/soccersim)

# Soccer Sim
### Run
1. Add a .env file with `APP_DB_USER`, `APP_DB_PASS`, `APP_DB_NAME`, and `APP_DB_ADDR` for a postgres database
1. Add Goose binary from https://github.com/pressly/goose
1. Run `goose postgres $APP_DB_ADDR up` in ROOT/data directory
1. Start server with `go run main.go`

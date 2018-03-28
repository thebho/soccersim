# Soccer Sim
### Run
1. Add a .env file with `APP_DB_USER`, `APP_DB_PASS`, `APP_DB_NAME`, and `APP_DB_ADDR` for a postgres database
1. Add Goose binary from https://github.com/pressly/goose
1. Run `goose postgres $APP_DB_ADDR up` in ROOT/data directory
1. Start server with `go run main.go`

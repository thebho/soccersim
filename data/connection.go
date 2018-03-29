package data

import (
	"os"

	"github.com/go-pg/pg"
)

// ConnectToDB initializes DB
func ConnectToDB() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     os.Getenv("APP_DB_USER"),
		Password: os.Getenv("APP_DB_PASS"),
		Addr:     os.Getenv("APP_DB_ADDR"),
		Database: os.Getenv("APP_DB_NAME"),
	})
}

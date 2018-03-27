package services

import (
	"os"

	"github.com/go-pg/pg"
)

var db *pg.DB

// InitDB init
func InitDB() {
	db = pg.Connect(&pg.Options{
		User:     os.Getenv("APP_DB_USER"),
		Password: os.Getenv("APP_DB_PASS"),
		Addr:     os.Getenv("APP_DB_ADDR"),
		Database: os.Getenv("APP_DB_NAME"),
	})
}

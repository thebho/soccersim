package data

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
	"github.com/subosito/gotenv"
)

func loadEnv() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}
}

func ConnectToDB() *pg.DB {
	fmt.Println(os.Getenv("APP_DB_USER"))
	return pg.Connect(&pg.Options{
		User:     os.Getenv("APP_DB_USER"),
		Password: os.Getenv("APP_DB_PASS"),
		Addr:     os.Getenv("APP_DB_ADDR"),
		Database: os.Getenv("APP_DB_NAME"),
	})
}
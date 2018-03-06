package main

import (
	"SoccerSim/model"
	"fmt"
	"os"

	"github.com/go-pg/pg"
	"github.com/subosito/gotenv"
)

var teams = []*model.Team{}

func init() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Connect to database
	db := pg.Connect(&pg.Options{
		User:     os.Getenv("APP_DB_USER"),
		Password: os.Getenv("APP_DB_PASS"),
		Addr:     os.Getenv("APP_DB_ADDR"),
		Database: os.Getenv("APP_DB_NAME"),
	})

	team := model.Team{Abv: "CRY"}
	err := db.Select(&team)
	if err != nil {
		panic(err)
	}
	fmt.Println(team)

}

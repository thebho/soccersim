package main

import (
	"SoccerSim/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var router *mux.Router

func main() {
	loadEnv()
	createRouter()
}

func loadEnv() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}
}

func createRouter() {
	router = mux.NewRouter()
	router.HandleFunc("/teams", controllers.GetTeams).Methods("GET")
	router.HandleFunc("/matches", controllers.ScheduleSeason).Methods("POST")
	router.HandleFunc("/season", controllers.ScheduleSeason).Methods("POST")
	matches := router.PathPrefix("/matches").Subrouter()

	// Start season
	// matches.HandleFunc("/", controllers.ScheduleSeason).Methods("POST")

	// Get weeks matches
	matches.HandleFunc("", controllers.GetWeeksMatches).Methods("GET")

	// Sim match week {seasonName, week, action}
	matches.HandleFunc("", controllers.SimWeeksMatches).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

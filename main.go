package main

import (
	"SoccerSim/handlers"
	"log"
	"net/http"

	"SoccerSim/util"

	"github.com/gorilla/mux"
)

var router *mux.Router
var soccerSim handlers.SoccerSim

func main() {
	util.LoadEnv()
	soccerSim = handlers.NewSoccerSim()
	createRouter()
}

func createRouter() {
	router = mux.NewRouter()
	router.HandleFunc("/teams", soccerSim.GetTeams).Methods("GET")
	router.HandleFunc("/season", soccerSim.ScheduleSeason).Methods("POST")

	matches := router.PathPrefix("/matches").Subrouter()

	// Get weeks matches
	matches.HandleFunc("", soccerSim.GetWeeksMatches).Methods("GET")

	// Sim match week {seasonName, week, action}
	matches.HandleFunc("", soccerSim.SimWeeksMatches).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

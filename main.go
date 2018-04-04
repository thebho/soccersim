package main

import (
	"log"
	"net/http"

	"github.com/thebho/soccersim/handlers"

	"github.com/thebho/soccersim/util"

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

	// Deprecated
	router.HandleFunc("/season", soccerSim.ScheduleSeason).Methods("POST")

	teams := router.PathPrefix("/teams").Subrouter()

	// Get all teams
	teams.HandleFunc("", soccerSim.GetWeeksMatches).Methods("GET")
	teams.HandleFunc("/{teamAbv}/seasons/{seasonKey}", soccerSim.GetTeamSeason).Methods("GET")

	router.HandleFunc("/teams", soccerSim.GetTeams).Methods("GET")

	seasons := router.PathPrefix("/seasons").Subrouter()
	seasons.HandleFunc("", soccerSim.ScheduleSeason).Methods("POST")
	seasons.HandleFunc("/{seasonKey}", soccerSim.GetSeasonTeams).Methods("GET")

	matches := router.PathPrefix("/matches").Subrouter()

	// Get weeks matches
	matches.HandleFunc("", soccerSim.GetWeeksMatches).Methods("GET")

	// Sim match week {seasonName, week, action}
	matches.HandleFunc("", soccerSim.SimWeeksMatches).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

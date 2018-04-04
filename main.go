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
	teams := router.PathPrefix("/teams").Subrouter()

	// Get all teams
	teams.HandleFunc("", soccerSim.GetWeeksMatches).Methods("GET")
	teams.HandleFunc("/{teamAbv}/seasons/{seasonKey}", soccerSim.GetTeamSeason).Methods("GET")
	teams.HandleFunc("/{teamAbv}/seasons", soccerSim.GetTeamSeasons).Methods("GET")

	router.HandleFunc("/teams", soccerSim.GetTeams).Methods("GET")
	router.HandleFunc("/season", soccerSim.ScheduleSeason).Methods("POST")

	matches := router.PathPrefix("/matches").Subrouter()

	// Get weeks matches
	matches.HandleFunc("", soccerSim.GetWeeksMatches).Methods("GET")

	// Sim match week {seasonName, week, action}
	matches.HandleFunc("", soccerSim.SimWeeksMatches).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

package main

import (
	"net/http"
	"os"

	"github.com/thebho/soccersim/handlers"

	"github.com/thebho/soccersim/util"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var router *mux.Router
var soccerSim handlers.SoccerSim

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Info("Starting Soccer Sim")
	util.LoadEnv()
	soccerSim = handlers.NewSoccerSim()
	createRouter()
}

func createRouter() {
	log.Info("Creating Router")
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

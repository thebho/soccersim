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
	router.HandleFunc("/matches", controllers.GetWeeksMatches).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetTeamSeason returns all teams from the database
func (s SoccerSim) GetTeamSeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	season := vars["seasonKey"]

	teamAbv := vars["teamAbv"]
	log.Printf("Getting team season for %s/%s", teamAbv, season)

	teamSeason := s.TeamService.GetTeamSeason(teamAbv, season)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(teamSeason)

}

// GetTeamSeason returns all teams from the database
func (s SoccerSim) GetTeamSeasons(w http.ResponseWriter, r *http.Request) {
	teamSeasons := s.TeamService.GetTeamSeasons()
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(teamSeasons)

}

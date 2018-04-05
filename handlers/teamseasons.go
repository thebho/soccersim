package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// GetTeamSeason returns all teams from the database
func (s SoccerSim) GetTeamSeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	season := vars["seasonKey"]

	teamAbv := vars["teamAbv"]
	log.Infof("Getting team season for %s/%s", teamAbv, season)

	teamSeason := s.TeamService.GetTeamSeason(teamAbv, season)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(teamSeason)

}

// GetSeasonTeams returns all teams from the database
func (s SoccerSim) GetSeasonTeams(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	season := vars["seasonKey"]
	log.Infof("Getting team season for %s", season)
	teamSeasons := s.TeamService.GetTeamSeasons(season)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(teamSeasons)

}

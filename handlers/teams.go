package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// GetTeams returns all teams from the database
func (s SoccerSim) GetTeams(w http.ResponseWriter, r *http.Request) {
	log.Info("Getting all teams")
	teams := s.TeamService.GetAllTeams()
	// util.WriteToFile(teams)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(teams)
}

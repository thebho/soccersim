package handlers_test

import (
	"SoccerSim/services"
	"SoccerSim/util"
	"encoding/json"
	"net/http"
	"testing"
)

// GetTeams returns all teams from the database
func TestGetTeams(t *testing.T) {
	req, err := http.NewRequest("GET", "/teams", nil)
	util.TestError(t, err)

	teamService := services.NewTeamService(s.db)
	teams := teamService.GetAllTeams()
	// util.WriteToFile(teams)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(teams)
}

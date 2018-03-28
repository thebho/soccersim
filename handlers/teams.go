package handlers

import (
	"SoccerSim/services"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTeams returns all teams from the database
func (s SoccerSim) GetTeams(w http.ResponseWriter, r *http.Request) {
	//TODO: Add logger
	fmt.Println("Getting all teams")
	teamService := services.NewTeamService(s.db)
	teams := teamService.GetAllTeams()
	// util.WriteToFile(teams)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(teams)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thebho/soccersim/util"
)

// GetTeams returns all teams from the database
func (s SoccerSim) GetTeams(w http.ResponseWriter, r *http.Request) {
	//TODO: Add logger
	fmt.Println("Getting all teams")
	season, err := util.ParseURL("season", r.URL.RequestURI())
	util.CheckError(err)
	teams := s.TeamService.GetAllTeams(season)
	// util.WriteToFile(teams)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(teams)
}

package handlers

import (
	"net/http"
	"soccersim/data"
	"soccersim/services/matches"
	"soccersim/services/teams"
)

// SoccerSim struct
type SoccerSim struct {
	TeamService  teams.TeamService
	MatchService matches.MatchService
}

// NewSoccerSim constructor
func NewSoccerSim() SoccerSim {
	db := data.NewPostgresDS()
	return SoccerSim{
		TeamService:  teams.NewTeamService(db),
		MatchService: matches.NewMatchService(db),
	}
}

func setReturnDefaults(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

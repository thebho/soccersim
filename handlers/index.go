package handlers

import (
	"net/http"

	"github.com/thebho/soccersim/data"
	"github.com/thebho/soccersim/services/matches"
	"github.com/thebho/soccersim/services/teams"
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

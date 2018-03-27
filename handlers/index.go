package handlers

import (
	"SoccerSim/data"
	"net/http"
)

// SoccerSim struct
type SoccerSim struct {
	db *data.PostgresDS
}

// NewSoccerSim constructor
func NewSoccerSim() SoccerSim {
	return SoccerSim{db: data.NewPostgresDS()}
}

func setReturnDefaults(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

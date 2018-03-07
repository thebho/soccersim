package controllers

import (
	"SoccerSim/model"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTeams returns all teams from the database
func GetTeams(w http.ResponseWriter, r *http.Request) {
	//TODO: Add logger
	fmt.Println("Getting all teams")
	var teams []model.Team
	err := controller.getDB().Model(&teams).Select()
	if err != nil {
		panic(err)
	}
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(teams)
}

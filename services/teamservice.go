package services

import "SoccerSim/model"

// GetAllTeams service
func GetAllTeams(dataStore model.TeamDataStore) []model.Team {
	return dataStore.GetTeams()
}

// GetTeam service
func GetTeam(dataStore model.TeamDataStore, teamABV string) model.Team {
	return dataStore.GetTeam(teamABV)
}

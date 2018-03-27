package services

import "SoccerSim/model"

// GetWeeksMatches service
func GetWeeksMatches(dataStore model.MatchesDataStore, season string, week int) []model.Match {
	return dataStore.GetMatches(season, week)
}

package services

import (
	"SoccerSim/model"
)

// MatchDataStore interface
type MatchDataStore interface {
	GetMatches(string, int) []model.Match
	GetTeam(string) model.Team
	UpdateObject(interface{})
}

// MatchServiceImp implements MatchService
type MatchServiceImp struct {
	dataStore MatchDataStore
}

// NewMatchService returns a MatchService with provided data store
func NewMatchService(dataStore MatchDataStore) MatchServiceImp {
	return MatchServiceImp{dataStore: dataStore}
}

// GetWeeksMatches service
func (m MatchServiceImp) GetWeeksMatches(season string, week int) []model.Match {
	return m.dataStore.GetMatches(season, week)
}

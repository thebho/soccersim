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

// Simulator describes a simulator
type Simulator interface {
	Sim(*model.Team, *model.Team, *model.Match)
}

// MatchServiceImp implements MatchService
type MatchServiceImp struct {
	dataStore      MatchDataStore
	matchSimulator Simulator
}

// NewMatchService returns a MatchService with provided data store
func NewMatchService(dataStore MatchDataStore) MatchServiceImp {
	return MatchServiceImp{dataStore: dataStore}
}

func NewMatchServiceSimulator(dataStore MatchDataStore,
	simulator Simulator) MatchServiceImp {
	return MatchServiceImp{dataStore: dataStore,
		matchSimulator: simulator}
}

// GetWeeksMatches service
func (m MatchServiceImp) GetWeeksMatches(season string,
	week int) []model.Match {
	return m.dataStore.GetMatches(season, week)
}

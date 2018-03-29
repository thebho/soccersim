package matches

import (
	"soccersim/model"
	"soccersim/simulator"
)

// MatchDataStore interface
type MatchDataStore interface {
	GetMatches(string, int) []model.Match
	GetTeam(string) model.Team
	UpdateObject(interface{})
	GetTeams() []model.Team
	SaveObject(interface{})
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
	return MatchServiceImp{dataStore: dataStore, matchSimulator: simulator.NewMatchSimulator()}
}

// GetWeeksMatches service
func (m MatchServiceImp) GetWeeksMatches(season string,
	week int) []model.Match {
	return m.dataStore.GetMatches(season, week)
}

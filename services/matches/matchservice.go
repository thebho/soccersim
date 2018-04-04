package matches

import (
	"github.com/thebho/soccersim/model"
	"github.com/thebho/soccersim/simulator"
)

// MatchDataStore interface
type MatchDataStore interface {
	GetMatches(string, int) []model.Match
	GetTeamSeason(string, string) model.TeamSeason
	UpdateObject(interface{})
	GetTeams() []model.Team
	SaveObject(interface{})
}

// Simulator describes a simulator
type Simulator interface {
	Sim(*model.TeamSeason, *model.TeamSeason, *model.Match)
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

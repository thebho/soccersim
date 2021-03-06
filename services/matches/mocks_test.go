package matches

import (
	"github.com/thebho/soccersim/model"
	"github.com/thebho/soccersim/util"

	"github.com/stretchr/testify/mock"
)

var matchA = model.NewMatch("ARS", "TOT", "MOCK", 1)
var matchB = model.NewMatch("WHA", "CHE", "MOCK", 1)

type MockMatchDataStore struct {
	mock.Mock
}

func (m *MockMatchDataStore) GetTeamSeason(teamAbv, season string) model.TeamSeason {
	_ = m.Called(teamAbv, season)
	return model.TeamSeason{Season: season, TeamID: teamAbv}
}
func (m *MockMatchDataStore) GetMatches(season string, matchWeek int) []model.Match {
	_ = m.Called(season, matchWeek)
	return []model.Match{matchA, matchB}
}
func (m *MockMatchDataStore) UpdateObject(object interface{}) {
	_ = m.Called(object)
	return
}
func (m *MockMatchDataStore) GetTeams() []model.Team {
	return util.GetTestTeams()
}
func (m *MockMatchDataStore) SaveObject(object interface{}) {
	_ = m.Called(object)
	return
}

type MockMatchSimulator struct {
	mock.Mock
}

func (m *MockMatchSimulator) Sim(homeTeamSeason, awayTeamSeason *model.TeamSeason, match *model.Match) {
	_ = m.Called(homeTeamSeason, awayTeamSeason, match)
	return
}

var teamA = model.Team{Name: "MockA", Abv: "MKA"}
var teamB = model.Team{Name: "MockB", Abv: "MKB"}
var teamC = model.Team{Name: "MockC", Abv: "MockC"}

func teamsHelper(numberOfTeams int) []model.Team {
	teams := []model.Team{teamA, teamB, teamC}
	return teams[0:numberOfTeams]
}

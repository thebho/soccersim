package teams

import (
	"errors"

	"github.com/thebho/soccersim/model"
	"github.com/thebho/soccersim/util"

	"github.com/stretchr/testify/mock"
)

var matchA = model.NewMatch("ARS", "TOT", "MOCK", 1)
var matchB = model.NewMatch("WHA", "CHE", "MOCK", 1)

type MockMatchDataStore struct {
	mock.Mock
}

func (m *MockMatchDataStore) GetTeam(teamAbv string) model.Team {
	_ = m.Called(teamAbv)
	return model.Team{Abv: teamAbv}
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

func (m *MockMatchSimulator) Sim(homeTeam, awayTeam *model.Team, match *model.Match) {
	_ = m.Called(homeTeam, awayTeam, match)
	return
}

type MockMatchesDataStore struct{}

func (m MockMatchesDataStore) GetMatches(season string, week int) []model.Match {
	return []model.Match{matchA, matchB}
}

var mockMatchDataStore = MockMatchesDataStore{}

var teamA = model.Team{Name: "MockA", Abv: "MKA"}
var teamB = model.Team{Name: "MockB", Abv: "MKB"}
var teamC = model.Team{Name: "MockC", Abv: "MockC"}
var teamASeasonA = model.TeamSeason{Season: "MockSeason", ID: 1, TeamID: teamA.Abv}
var teamASeasonJoinA = model.NewTeamSeasonJoin(teamASeasonA.ID, teamA.Abv, teamA.Name, teamASeasonA.Season)

type MockTeamDataStore struct{}

func (m MockTeamDataStore) GetTeam(teamABV string) model.Team {
	if teamABV == teamA.Abv {
		return teamA
	}
	panic(errors.New(""))
}

func (m MockTeamDataStore) GetTeamSeason(teamABV, season string) model.TeamSeason {
	if teamABV == teamA.Abv {
		return teamASeasonA
	}
	panic(errors.New(""))
}
func (m MockTeamDataStore) GetTeamSeasons(season string) []model.TeamSeasonJoin {
	return nil
}
func (m MockTeamDataStore) GetTeams() []model.Team {
	return []model.Team{teamA, teamB}
}

var mockTeamDataStore = MockTeamDataStore{}

type MockSchedulerDataStore struct{}

func (m MockSchedulerDataStore) SaveObject(object interface{}) {}
func (m MockSchedulerDataStore) GetTeams() []model.Team {
	return util.GetTestTeams()
}

var mockSchedulerDataStore = MockSchedulerDataStore{}

func teamsHelper(numberOfTeams int) []model.Team {
	teams := []model.Team{teamA, teamB, teamC}
	return teams[0:numberOfTeams]
}

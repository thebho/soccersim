package handlers_test

import (
	"github.com/thebho/soccersim/model"
	"github.com/thebho/soccersim/util"

	"github.com/stretchr/testify/mock"
)

type MockTeamService struct {
	mock.Mock
}

func (m *MockTeamService) GetAllTeams() []model.Team {
	m.Called()
	return util.GetTestTeamsPath("../util/teams.txt")
}
func (m *MockTeamService) GetTeam(teamAbv string) model.Team {
	m.Called(teamAbv)
	return model.Team{}
}
func (m *MockTeamService) GetTeamSeason(teamAbv, season string) model.TeamSeason {
	m.Called(teamAbv, season)
	return model.TeamSeason{}
}
func (m *MockTeamService) GetTeamSeasons() []model.TeamSeason {
	m.Called()
	return nil
}

type MockMatchService struct {
	mock.Mock
}

func (m *MockMatchService) GetWeeksMatches(season string, week int) []model.Match {
	m.Called(season, week)
	return []model.Match{model.Match{}, model.Match{}}
}
func (m *MockMatchService) SimMatchWeek(season string, week int) {
	m.Called(season, week)
	return
}
func (m *MockMatchService) ScheduleSeason(seasonName string) []model.Match {
	m.Called(seasonName)
	return []model.Match{model.Match{}}
}

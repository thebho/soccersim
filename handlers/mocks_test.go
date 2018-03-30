package handlers_test

import (
	"github.com/thebho/soccersim/model"
	"github.com/thebho/soccersim/util"

	"github.com/stretchr/testify/mock"
)

type MockTeamService struct {
	mock.Mock
}

func (m *MockTeamService) GetAllTeams(season string) []model.Team {
	m.Called(season)
	return util.GetTestTeamsPath("../util/teams.txt")
}
func (m *MockTeamService) GetTeam(teamAbv string) model.Team {
	m.Called(teamAbv)
	return model.Team{}
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

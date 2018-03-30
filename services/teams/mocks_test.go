package teams

import (
	"github.com/thebho/soccersim/model"

	"github.com/stretchr/testify/mock"
)

var matchA = model.NewMatch("ARS", "TOT", "MOCK", 1)
var matchB = model.NewMatch("WHA", "CHE", "MOCK", 1)

type MockTeamDataStore struct {
	mock.Mock
}

func (m *MockTeamDataStore) GetTeam(teamAbv string) model.Team {
	_ = m.Called(teamAbv)
	return teamA
}

func (m *MockTeamDataStore) GetTeams() []model.Team {
	return []model.Team{teamA, teamB, teamC}
}

func (m *MockTeamDataStore) GetTeamsBySeason(season string) []model.Team {
	if season == "MKSeason" {
		return []model.Team{teamA, teamB}
	}
	return nil
}

var teamA = model.Team{Name: "MockA", Abv: "MKA", Season: "MKSeason"}
var teamB = model.Team{Name: "MockB", Abv: "MKB", Season: "MKSeason"}
var teamC = model.Team{Name: "MockC", Abv: "MockC"}

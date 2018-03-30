package teams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// var teamService = NewTeamService(mockTeamDataStore)

func TestGetAllTeamsNoSeason(t *testing.T) {
	mockTDS := new(MockTeamDataStore)
	mockTDS.On("GetAllTeams", "")
	teamService := NewTeamService(mockTDS)
	teams := teamService.GetAllTeams("")
	assert.Equal(t, teamA.Abv, teams[0].Abv)
	assert.Equal(t, teamB.Name, teams[1].Name)
	assert.Equal(t, 3, len(teams))
}

func TestGetAllTeamsSeason(t *testing.T) {
	mockTDS := new(MockTeamDataStore)
	mockTDS.On("GetAllTeams", "MKSeason")
	teamService := NewTeamService(mockTDS)
	teams := teamService.GetAllTeams("MKSeason")
	assert.Equal(t, teamA.Abv, teams[0].Abv)
	assert.Equal(t, teamB.Name, teams[1].Name)
	assert.Equal(t, 2, len(teams))
}

func TestGetTeam(t *testing.T) {
	mockTDS := new(MockTeamDataStore)
	mockTDS.On("GetTeam", teamA.Abv)
	teamService := NewTeamService(mockTDS)
	team := teamService.GetTeam("MKA")
	assert.Equal(t, teamA.Abv, team.Abv)
	assert.Equal(t, teamA.Name, team.Name)
}

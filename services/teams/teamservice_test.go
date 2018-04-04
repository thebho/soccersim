package teams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var teamService = NewTeamService(mockTeamDataStore)

func TestGetAllTeams(t *testing.T) {
	teams := teamService.GetAllTeams()
	assert.Equal(t, teamA.Abv, teams[0].Abv)
	assert.Equal(t, teamB.Name, teams[1].Name)
}

func TestGetTeam(t *testing.T) {
	team := teamService.GetTeam("MKA")
	assert.Equal(t, teamA.Abv, team.Abv)
	assert.Equal(t, teamA.Name, team.Name)
}

func TestGetTeamSeason(t *testing.T) {
	teamSeason := teamService.GetTeamSeason("MKA", "MockSeason")
	assert.Equal(t, teamA.Abv, teamSeason.Abv)
	assert.Equal(t, teamASeasonA.ID, teamSeason.ID)
}

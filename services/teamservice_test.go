package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTeams(t *testing.T) {
	teams := GetAllTeams(mockTeamDataStore)
	assert.Equal(t, teamA.Abv, teams[0].Abv)
	assert.Equal(t, teamB.Name, teams[1].Name)
}

func TestGetTeam(t *testing.T) {
	team := GetTeam(mockTeamDataStore, "MKA")
	assert.Equal(t, teamA.Abv, team.Abv)
	assert.Equal(t, teamA.Name, team.Name)
}

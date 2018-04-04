package simulator

import (
	"testing"

	"github.com/thebho/soccersim/model"

	"github.com/stretchr/testify/assert"
)

func TestMatchPlayed(t *testing.T) {
	matchSimulator := testMatchSimulator([]int{1}, []int{0})
	teamASeason := model.TeamSeason{}
	teamBSeason := model.TeamSeason{}
	match := &model.Match{}
	matchSimulator.Sim(&teamASeason, &teamBSeason, match)
	assert.Equal(t, true, match.Played)
}

func TestMatchSimulatorFixedResults(t *testing.T) {
	matchSimulator := testMatchSimulator([]int{1}, []int{0})
	teamA := &model.TeamSeason{}
	teamB := &model.TeamSeason{}
	match := &model.Match{}
	matchSimulator.Sim(teamA, teamB, match)
	assert.Equal(t, 1, teamA.GamesWon)
	assert.Equal(t, 0, teamA.GamesLost)
	assert.Equal(t, 0, teamA.GamesDrawn)
	assert.Equal(t, 1, teamA.GoalsScored)
	assert.Equal(t, 0, teamA.GoalsAllowed)
	assert.Equal(t, 1, teamB.GamesLost)
	assert.Equal(t, 0, teamB.GamesWon)
	assert.Equal(t, 0, teamB.GamesDrawn)
	assert.Equal(t, 1, teamB.GoalsAllowed)
	assert.Equal(t, 0, teamB.GoalsScored)
	assert.Equal(t, true, match.Played)
}

func TestMatchSimulatorDraw(t *testing.T) {
	matchSimulator := testMatchSimulator([]int{1}, []int{1})

	teamASeason := model.TeamSeason{}
	teamBSeason := model.TeamSeason{}
	match := &model.Match{}
	matchSimulator.Sim(&teamASeason, &teamBSeason, match)
	assert.Equal(t, 0, teamASeason.GamesWon)
	assert.Equal(t, 0, teamASeason.GamesLost)
	assert.Equal(t, 1, teamASeason.GamesDrawn)
	assert.Equal(t, 1, teamASeason.GoalsScored)
	assert.Equal(t, 1, teamASeason.GoalsAllowed)
	assert.Equal(t, 0, teamBSeason.GamesLost)
	assert.Equal(t, 0, teamBSeason.GamesWon)
	assert.Equal(t, 1, teamBSeason.GamesDrawn)
	assert.Equal(t, 1, teamBSeason.GoalsAllowed)
	assert.Equal(t, 1, teamBSeason.GoalsScored)
}

func TestMatchSimRandom(t *testing.T) {
	matchSimulator := NewMatchSimulator()

	var teamASeason model.TeamSeason
	var teamBSeason model.TeamSeason
	match := &model.Match{HomeTeam: "A", AwayTeam: "B"}
	matchSimulator.Sim(&teamASeason, &teamBSeason, match)
	assert.NotNil(t, match.AwayTeamGoals)
	assert.NotNil(t, match.HomeTeamGoals)
}

package simulator

import (
	"soccersim/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchPlayed(t *testing.T) {
	matchSimulator := testMatchSimulator([]int{1}, []int{0})
	teamA := &model.Team{}
	teamB := &model.Team{}
	match := &model.Match{}
	matchSimulator.Sim(teamA, teamB, match)
	assert.Equal(t, true, match.Played)
}

func TestMatchSimulatorFixedResults(t *testing.T) {
	matchSimulator := testMatchSimulator([]int{1}, []int{0})
	teamA := &model.Team{}
	teamB := &model.Team{}
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
	teamA := &model.Team{}
	teamB := &model.Team{}
	match := &model.Match{}
	matchSimulator.Sim(teamA, teamB, match)
	assert.Equal(t, 0, teamA.GamesWon)
	assert.Equal(t, 0, teamA.GamesLost)
	assert.Equal(t, 1, teamA.GamesDrawn)
	assert.Equal(t, 1, teamA.GoalsScored)
	assert.Equal(t, 1, teamA.GoalsAllowed)
	assert.Equal(t, 0, teamB.GamesLost)
	assert.Equal(t, 0, teamB.GamesWon)
	assert.Equal(t, 1, teamB.GamesDrawn)
	assert.Equal(t, 1, teamB.GoalsAllowed)
	assert.Equal(t, 1, teamB.GoalsScored)
}

func TestMatchSimRandom(t *testing.T) {
	matchSimulator := NewMatchSimulator()
	teamA := &model.Team{Abv: "A"}
	teamB := &model.Team{Abv: "B"}
	match := &model.Match{HomeTeam: "A", AwayTeam: "B"}
	matchSimulator.Sim(teamA, teamB, match)
	assert.NotNil(t, match.AwayTeamGoals)
	assert.NotNil(t, match.HomeTeamGoals)
}

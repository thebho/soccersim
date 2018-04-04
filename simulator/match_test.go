package simulator

import (
	"testing"

	"github.com/thebho/soccersim/model"

	"github.com/stretchr/testify/assert"
)

func TestMatchPlayed(t *testing.T) {
	matchSimulator := testMatchSimulator([]int{1}, []int{0})
	teamA := &model.Team{}
	teamB := &model.Team{}
	teamASeason := model.NewTeamSeasonJoin(1, teamA.Abv, teamA.Name, "MockSeason")
	teamBSeason := model.NewTeamSeasonJoin(2, teamB.Abv, teamB.Name, "MockSeason")
	match := &model.Match{}
	matchSimulator.Sim(&teamASeason, &teamBSeason, match)
	assert.Equal(t, true, match.Played)
}

func TestMatchSimulatorFixedResults(t *testing.T) {
	matchSimulator := testMatchSimulator([]int{1}, []int{0})
	teamA := &model.TeamSeasonJoin{}
	teamB := &model.TeamSeasonJoin{}
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
	teamASeason := model.NewTeamSeasonJoin(1, teamA.Abv, teamA.Name, "MockSeason")
	teamBSeason := model.NewTeamSeasonJoin(2, teamB.Abv, teamB.Name, "MockSeason")
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
	teamA := &model.Team{Abv: "A"}
	teamB := &model.Team{Abv: "B"}
	teamASeason := model.NewTeamSeasonJoin(1, teamA.Abv, teamA.Name, "MockSeason")
	teamBSeason := model.NewTeamSeasonJoin(2, teamB.Abv, teamB.Name, "MockSeason")
	match := &model.Match{HomeTeam: "A", AwayTeam: "B"}
	matchSimulator.Sim(&teamASeason, &teamBSeason, match)
	assert.NotNil(t, match.AwayTeamGoals)
	assert.NotNil(t, match.HomeTeamGoals)
}

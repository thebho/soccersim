package services

import (
	"SoccerSim/model"
	"SoccerSim/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testSeasonConstant = "Test Season"

func teamsHelper(numberOfTeams int) []model.Team {
	teamA := model.Team{Abv: "MockA"}
	teamB := model.Team{Abv: "MockB"}
	teamC := model.Team{Abv: "MockC"}
	teams := []model.Team{teamA, teamB, teamC}
	return teams[0:numberOfTeams]
}

//
// func TestScheduleSeason(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip()
// 	}
// 	mockTeams := util.GetTestTeams()
// 	scheduler := SchedulerImp{}
// 	matchUps := scheduler.ScheduleSeason(testSeasonConstant)
// 	assert.Equal(t, 380, len(matchUps))
// 	assert.Equal(t, testSeasonConstant, matchUps[0].Season)
// }

func TestCreateMatchUps(t *testing.T) {
	teams := teamsHelper(2)
	matches := createMatchUps(teams)
	assert.Equal(t, teams[0].Abv, matches[0].HomeTeam, "A Team is home")
	assert.Equal(t, teams[1].Abv, matches[0].AwayTeam, "B Team is away")
	assert.Equal(t, teams[1].Abv, matches[1].HomeTeam, "B Team is home")
	assert.Equal(t, teams[0].Abv, matches[1].AwayTeam, "A Team is away")
}

func TestScheduledMatchUps(t *testing.T) {
	teams := util.GetTestTeams()
	matchups := scheduleMatchUps(teams, "MockSeasonName")
	assert.Equal(t, 380, len(matchups))
	// assert.NotEqual(t, 0, scheduledMatches[0].MatchWeek, "Match 1 has a MatchWeek")

	// matches per team
	ars := 0
	for _, match := range matchups {
		if match.HomeTeam == "ARS" || match.AwayTeam == "ARS" {
			ars++
		}
	}
	assert.Equal(t, 38, ars)

	// matches per week
	week10 := 0
	for _, match := range matchups {
		if match.MatchWeek == 10 {
			week10++

		}
	}
	assert.Equal(t, 10, week10)
}

func TestPosition(t *testing.T) {
	assert.Equal(t, 3, position(3, 0, 20))
	assert.Equal(t, 4, position(3, 1, 20))
	assert.Equal(t, 2, position(3, 18, 20))
	assert.Equal(t, 13, position(3, 10, 20))
	assert.Equal(t, 19, position(19, 0, 20))
	assert.Equal(t, 1, position(19, 1, 20))
	assert.Equal(t, 18, position(19, 18, 20))
	assert.Equal(t, 10, position(19, 10, 20))
}

func TestCreateTeamMatrix(t *testing.T) {
	mockTeams := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	weeks := createTeamMatrix(mockTeams)
	assert.Equal(t, twoTeams{0, 1}, weeks[1][0])
	assert.Equal(t, twoTeams{8, 9}, weeks[18][9])
}

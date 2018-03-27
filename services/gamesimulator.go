package services

import (
	"SoccerSim/model"

	rand "github.com/thebho/random-tools"
)

var homeTeamGoalsArray = []int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 4, 5, 6}
var awayTeamGoalsArray = []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 3, 3, 4}

// SimGame takes a home "Team" and away "Team" and prints results
func SimGame(homeTeam, awayTeam *model.Team, match *model.Match) {
	//TODO: Add variance in the weighted arrays (if a team is +1 form, only use indexes 1:len(array), if they are -2 only use indexes 0:len(array)-2...)
	homeTeamGoals := rand.RandIntFromArray(homeTeamGoalsArray)
	awayTeamGoals := rand.RandIntFromArray(awayTeamGoalsArray)

	homeTeam.AddResult(homeTeamGoals, awayTeamGoals)
	awayTeam.AddResult(awayTeamGoals, homeTeamGoals)
	match.HomeTeamGoals = homeTeamGoals
	match.AwayTeamGoals = awayTeamGoals
	match.Played = true
}

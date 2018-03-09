package services

import (
	"SoccerSim/model"
	"fmt"

	"github.com/go-pg/pg"
	randtools "github.com/thebho/random-tools"
)

// Scheduler takes a season name and an slice of teams and returns and stores the season in the database
type Scheduler interface {
	ScheduleSeason(string, []model.Team)
}

// SchedulerImp is an implementation of Scheduler
type SchedulerImp struct {
	db *pg.DB
}

// ScheduleSeason schedules a new season from an slice of teams
func (s SchedulerImp) ScheduleSeason(seasonName string, teams []model.Team) {
	// if s.db == nil {
	// 	s.db = data.ConnectToDB()
	// }
	fmt.Printf("Scheduling season: %s\n", seasonName)
	_ = createMatchUps(teams)
	fmt.Println("Matches Created")
}

func createMatchUps(teams []model.Team) []model.Match {
	var matches []model.Match
	for home := range teams {
		for away := range teams {
			if teams[home].Abv == teams[away].Abv {
				continue
			}
			var match = model.Match{
				HomeTeam: teams[home].Abv,
				AwayTeam: teams[away].Abv,
			}
			matches = append(matches, match)
		}
	}
	return matches
}

type MatchUpHelper struct {
	teamA int
	teamB int
}

func createNumberOfTeamsSlice(numberOfTeams int) []int {
	var allTeams = make([]int, numberOfTeams, numberOfTeams)
	for i := 0; i < numberOfTeams; i++ {
		allTeams[i] = i
	}
	return allTeams
}

func createMatchUpAlgorithm(numberOfTeams int) {
	// matchUpMap := make(map[string]bool)
	// var weeks = make(map[int]MatchUpHelper)
	for i := 0; i < numberOfTeams/2; i++ {
		var teamsAvailable = createNumberOfTeamsSlice(numberOfTeams)
		fmt.Println(teamsAvailable)

	}
}

// func createMatchUpMatrix(numberOfTeams int) {
// 	teamsSlice := createNumberOfTeamsSlice(numberOfTeams)
// }

type twoTeams struct {
	a int
	b int
}

// 1  20
// 2  19
// 3  18
// 4  17
// 5  16
// 6  15
// 7  14
// 8  13
// 9  12
// 10 11
func newTwoTeams(a, b int) twoTeams { return twoTeams{a: a, b: b} }

func position(pos, offset, totalTeams int) int {
	if pos+offset < totalTeams {
		return pos + offset
	}
	return ((pos + offset) - totalTeams) + 1
}

func pos(a, b int) int {
	return position(a, b, 20)
}

func createTeamMatrix(teams []int) [][]twoTeams {
	var teamMatrix = make([][]twoTeams, 19, 19)
	for i := 0; i < 19; i++ {
		teamMatrix[i] = []twoTeams{newTwoTeams(teams[0], teams[pos(19, i)]),
			newTwoTeams(teams[pos(1, i)], teams[pos(18, i)]),
			newTwoTeams(teams[pos(2, i)], teams[pos(17, i)]),
			newTwoTeams(teams[pos(3, i)], teams[pos(16, i)]),
			newTwoTeams(teams[pos(4, i)], teams[pos(15, i)]),
			newTwoTeams(teams[pos(5, i)], teams[pos(14, i)]),
			newTwoTeams(teams[pos(6, i)], teams[pos(13, i)]),
			newTwoTeams(teams[pos(7, i)], teams[pos(12, i)]),
			newTwoTeams(teams[pos(8, i)], teams[pos(11, i)]),
			newTwoTeams(teams[pos(9, i)], teams[pos(10, i)]),
		}
	}
	return teamMatrix
}

func scheduleMatchUps(unscheduledMatches []model.Match, teams []model.Team) []model.Match {
	// _ = createSchedulerHelperArray(teams)
	var scheduledMatches []model.Match
	intArray := randtools.Array(len(teams))
	fmt.Printf("Teams array: %v\n", intArray)
	return scheduledMatches
}

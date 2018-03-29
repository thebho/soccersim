package matches

import (
	"fmt"

	"github.com/thebho/soccersim/model"

	randtools "github.com/thebho/random-tools"
)

//SchedulerDataStore interface
type SchedulerDataStore interface {
	GetTeams() []model.Team
	SaveObject(interface{})
}

// ScheduleSeason schedules a new season from an slice of teams
func (m MatchServiceImp) ScheduleSeason(seasonName string) []model.Match {
	fmt.Printf("Scheduling season: %s\n", seasonName)
	teams := m.dataStore.GetTeams()
	matches := m.scheduleMatchUps(teams, seasonName)
	m.insertIntoDB(matches)
	return matches
}

func (m MatchServiceImp) insertIntoDB(matches []model.Match) {
	for i := range matches {
		m.dataStore.SaveObject(&matches[i])
	}

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

type twoTeams struct {
	a int
	b int
}

func newTwoTeams(a, b int) twoTeams { return twoTeams{a: a, b: b} }

func position(pos, offset, totalTeams int) int {
	if pos+offset < totalTeams {
		return pos + offset
	}
	return ((pos + offset) - totalTeams) + 1
}

func (m MatchServiceImp) scheduleMatchUps(teams []model.Team, name string) []model.Match {
	// _ = createSchedulerHelperArray(teams)
	var scheduledMatches []model.Match
	teamMatrixHalf := createTeamMatrix(randtools.Array(len(teams)))
	for i, week := range teamMatrixHalf {
		weekIndex := i + 1
		for j := (weekIndex % 2); j < 10; j += 2 {
			scheduledMatches = append(scheduledMatches, model.NewMatch(teams[week[j].a].Abv, teams[week[j].b].Abv, name, weekIndex))
			scheduledMatches = append(scheduledMatches, model.NewMatch(teams[week[j].b].Abv, teams[week[j].a].Abv, name, weekIndex+19))
		}

		for j := ((weekIndex + 1) % 2); j < 10; j += 2 {
			scheduledMatches = append(scheduledMatches, model.NewMatch(teams[week[j].b].Abv, teams[week[j].a].Abv, name, weekIndex))
			scheduledMatches = append(scheduledMatches, model.NewMatch(teams[week[j].a].Abv, teams[week[j].b].Abv, name, weekIndex+19))
		}
	}
	return scheduledMatches
}

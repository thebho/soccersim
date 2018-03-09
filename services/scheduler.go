package services

import (
	"SoccerSim/model"
	"fmt"

	"github.com/go-pg/pg"
	rand "github.com/thebho/random-tools"
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

type scheduleHelperTeam struct {
	team               model.Team
	matchWeekScheduled bool
	matches            []model.Match
}

func createSchedulerHelperArray(teams []model.Team) map[string]scheduleHelperTeam {
	var scheduleHelpers = make(map[string]scheduleHelperTeam)
	for i := range teams {
		scheduleHelpers[teams[i].Abv] = scheduleHelperTeam{team: teams[i]}
	}
	return scheduleHelpers
}

func scheduleMatchUps(unscheduledMatches []model.Match, teams []model.Team) []model.Match {
	schedulerHelpers := createSchedulerHelperArray(teams)
	var scheduledMatches []model.Match
	rand.RandIndex(count)

	return scheduledMatches
}

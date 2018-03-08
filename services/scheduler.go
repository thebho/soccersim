package services

import (
	"SoccerSim/model"
	"fmt"

	"github.com/go-pg/pg"
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

}

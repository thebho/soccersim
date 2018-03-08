package services

import (
	"SoccerSim/util"
	"fmt"
	"testing"
)

func TestScheduleSeason(t *testing.T) {
	fmt.Println("Testing Schedule Season")
	mockTeams := util.GetTestTeams()
	scheduler := SchedulerImp{}
	scheduler.ScheduleSeason("Test Season", mockTeams)

}

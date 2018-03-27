package services

import (
	"SoccerSim/model"
	"SoccerSim/util"
	"errors"
)

var matchA = model.NewMatch("ARS", "TOT", "MOCK", 1)
var matchB = model.NewMatch("WHA", "CHE", "MOCK", 1)

type MockMatchesDataStore struct{}

func (m MockMatchesDataStore) GetMatches(season string, week int) []model.Match {
	return []model.Match{matchA, matchB}
}

var mockMatchDataStore = MockMatchesDataStore{}

var teamA = model.Team{Name: "MockA", Abv: "MKA"}
var teamB = model.Team{Name: "MockB", Abv: "MKB"}
var teamC = model.Team{Name: "MockC", Abv: "MockC"}

type MockTeamDataStore struct{}

func (m MockTeamDataStore) GetTeam(teamABV string) model.Team {
	if teamABV == teamA.Abv {
		return teamA
	}
	panic(errors.New(""))
}
func (m MockTeamDataStore) GetTeams() []model.Team {
	return []model.Team{teamA, teamB}
}

var mockTeamDataStore = MockTeamDataStore{}

type MockSchedulerDataStore struct{}

func (m MockSchedulerDataStore) SaveObject(object interface{}) {}
func (m MockSchedulerDataStore) GetTeams() []model.Team {
	return util.GetTestTeams()
}

var mockSchedulerDataStore = MockSchedulerDataStore{}

func teamsHelper(numberOfTeams int) []model.Team {
	teams := []model.Team{teamA, teamB, teamC}
	return teams[0:numberOfTeams]
}

type MockMatchWeekDataStore struct{}

var matchSimA = model.NewMatch(teamA.Abv, teamB.Abv, "MOCK", 1)

func (m MockMatchWeekDataStore) GetTeam(teamABV string) model.Team {
	if teamABV == teamA.Abv {
		return teamA
	} else if teamABV == teamB.Abv {
		return teamB
	} else {
		panic(errors.New("Unknown teamABV"))
	}
}
func (m MockMatchWeekDataStore) GetWeeksMatches(season string, week int) []model.Match {
	return []model.Match{matchSimA}
}

var mockMatchWeekDataStore = MockMatchWeekDataStore{}

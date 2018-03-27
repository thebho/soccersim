package services

import (
	"SoccerSim/model"
	"fmt"
)

// MatchWeekDataStore interface
type MatchWeekDataStore interface {
	GetMatches(string, int) []model.Match
	GetTeam(string) model.Team
	UpdateObject(interface{})
}

type simulator struct {
	dataStore MatchWeekDataStore
}

//SimMatchWeek takes a season and a matchWeek and simulates all the matches that week
/*
1. Pulls the matches from that week from the DB
2. Runs the game simulator to get the game score from the 2 teams
3. Updates the match with the result in the db
4. Updates the teams records and goals
*/
func SimMatchWeek(dataStore MatchWeekDataStore, season string, matchWeek int) {
	fmt.Printf("Simming match week: %d in season: %s\n", matchWeek, season)
	weekSimulator := simulator{dataStore: dataStore}
	matches := dataStore.GetMatches(season, matchWeek)
	for _, match := range matches {
		weekSimulator.simMatch(match)
	}
}

func (s simulator) simMatch(match model.Match) {
	if match.Played {
		fmt.Printf("Skipping previously played match: %d\n", match.ID)
		return
	}
	homeTeam := s.dataStore.GetTeam(match.HomeTeam)
	awayTeam := s.dataStore.GetTeam(match.AwayTeam)

	SimGame(&homeTeam, &awayTeam, &match)
	fmt.Println(homeTeam)
	s.dataStore.UpdateObject(&homeTeam)
	s.dataStore.UpdateObject(&awayTeam)
	s.dataStore.UpdateObject(&match)
}

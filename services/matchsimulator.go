package services

import (
	"SoccerSim/model"
	"fmt"
)

type MatchWeekDataStore interface {
	GetWeeksMatches(string, int) []model.Match
	GetTeam(string) model.Team
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
	matches := dataStore.GetWeeksMatches(season, matchWeek)
	for _, match := range matches {
		weekSimulator.simMatch(match)
	}
}

func (s simulator) simMatch(match model.Match) {
	homeTeam := s.dataStore.GetTeam(match.HomeTeam)
	awayTeam := s.dataStore.GetTeam(match.AwayTeam)

	/* TODO: Rework sim game to...
	pull in match and database
	confirm match hasn't been played already
	find teams from the DB
	sim game
	update match in DB
	update Teams in DB
	*/
	SimGame(&homeTeam, &awayTeam)
	fmt.Printf("Home Team: %v, Away Team: %v\n", homeTeam, awayTeam)

}

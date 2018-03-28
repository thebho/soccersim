package services

import (
	"SoccerSim/model"
	"fmt"
)

//SimMatchWeek takes a season and a matchWeek and simulates all the matches that week
/*
1. Pulls the matches from that week from the DB
2. Runs the game simulator to get the game score from the 2 teams
3. Updates the match with the result in the db
4. Updates the teams records and goals
*/
func (m MatchServiceImp) SimMatchWeek(season string, matchWeek int) {
	fmt.Printf("Simming match week: %d in season: %s\n", matchWeek, season)
	matches := m.dataStore.GetMatches(season, matchWeek)
	for _, match := range matches {
		m.simMatch(match)
	}
}

func (m MatchServiceImp) simMatch(match model.Match) {
	if match.Played {
		fmt.Printf("Skipping previously played match: %d\n", match.ID)
		return
	}
	homeTeam := m.dataStore.GetTeam(match.HomeTeam)
	awayTeam := m.dataStore.GetTeam(match.AwayTeam)

	SimGame(&homeTeam, &awayTeam, &match)
	fmt.Println(homeTeam)
	m.dataStore.UpdateObject(&homeTeam)
	m.dataStore.UpdateObject(&awayTeam)
	m.dataStore.UpdateObject(&match)
}

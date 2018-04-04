package matches

import (
	"fmt"

	"github.com/thebho/soccersim/model"
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
		m.simMatch(match, season)
	}
}

func (m MatchServiceImp) simMatch(match model.Match, season string) {
	if match.Played {
		fmt.Printf("Skipping previously played match: %d\n", match.ID)
		return
	}
	homeTeamSeason := m.dataStore.GetTeamSeason(match.HomeTeam, season)
	awayTeamSeason := m.dataStore.GetTeamSeason(match.AwayTeam, season)
	m.matchSimulator.Sim(&homeTeamSeason, &awayTeamSeason, &match)
	m.dataStore.UpdateObject(&homeTeamSeason)
	m.dataStore.UpdateObject(&awayTeamSeason)
	m.dataStore.UpdateObject(&match)
}

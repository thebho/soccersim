package model

// TeamSeason data model
type TeamSeason struct {
	tableName    struct{} `sql:"team_season, alias:team_seasons"`
	ID           int
	Season       string
	TeamID       string
	GamesWon     int
	GamesLost    int
	GamesDrawn   int
	GoalsScored  int
	GoalsAllowed int
}

// AddResult updates a teams record and Goals
func (t *TeamSeason) AddResult(goalsFor, goalsAgainst int) {
	t.GoalsScored += goalsFor
	t.GoalsAllowed += goalsAgainst
	if goalsFor > goalsAgainst {
		// fmt.Printf("%s Wins! %d-%d\n", t.Name, goalsFor, goalsAgainst)
		t.GamesWon++
	} else if goalsFor == goalsAgainst {
		t.GamesDrawn++
		// fmt.Printf("Draw %d-%d\n", goalsFor, goalsAgainst)
	} else {
		t.GamesLost++
		// fmt.Printf("%s Loses %d-%d\n", t.Name, goalsAgainst, goalsFor)
	}
}

// TotalPoints converts record to league points
func (t TeamSeason) TotalPoints() int {
	return t.GamesWon*3 + t.GamesDrawn
}

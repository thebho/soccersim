package model

// Team data model
type Team struct {
	Abv          string `sql:",pk"`
	Name         string
	GamesWon     int
	GamesLost    int
	GamesDrawn   int
	GoalsScored  int
	GoalsAllowed int
}

// AddResult updates a teams record and Goals
func (t *Team) AddResult(goalsFor, goalsAgainst int) {
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
func (t Team) TotalPoints() int {
	return t.GamesWon*3 + t.GamesDrawn
}

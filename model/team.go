package model

type Team struct {
	name         string
	gamesWon     int
	gamesLost    int
	gamesDrawn   int
	goalsScored  int
	goalsAllowed int
}

func NewTeam(teamName string) *Team {
	return &Team{
		name: teamName,
	}
}
func (t Team) Name() string {
	return t.name
}

// AddResult updates a teams record and goals
func (t *Team) AddResult(goalsFor, goalsAgainst int) {
	t.goalsScored += goalsFor
	t.goalsAllowed += goalsAgainst
	if goalsFor > goalsAgainst {
		// fmt.Printf("%s Wins! %d-%d\n", t.name, goalsFor, goalsAgainst)
		t.gamesWon++
	} else if goalsFor == goalsAgainst {
		t.gamesDrawn++
		// fmt.Printf("Draw %d-%d\n", goalsFor, goalsAgainst)
	} else {
		t.gamesLost++
		// fmt.Printf("%s Loses %d-%d\n", t.name, goalsAgainst, goalsFor)
	}
}

// TotalPoints converts record to league points
func (t Team) TotalPoints() int {
	return t.gamesWon*3 + t.gamesDrawn
}

// GetWins returns number of team wins
func (t Team) GetWins() int {
	return t.gamesWon
}

// GetDraws returns number of team draws
func (t Team) GetDraws() int {
	return t.gamesDrawn
}

// GetLoses returns number of team losses
func (t Team) GetLoses() int {
	return t.gamesLost
}

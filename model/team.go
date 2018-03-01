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

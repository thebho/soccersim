package model

type Team struct {
	name         string
	gamesWon     int
	gamesLost    int
	gamesDrawn   int
	goalsScored  int
	goalsAllowed int
}

func NewTeam(teamName string) Team {
	return Team{
		name: teamName,
	}
}
func (t Team) Name() string {
	return t.name
}

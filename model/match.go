package model

//Match data model
type Match struct {
	HomeTeam      string
	AwayTeam      string
	HomeTeamGoals int
	AwayTeamGoals int
	MatchWeek     int
	Season        string
}

package model

//Match data model
type Match struct {
	tableName     struct{} `sql:"matches, alias:match"`
	ID            int
	HomeTeam      string
	AwayTeam      string
	HomeTeamGoals int
	AwayTeamGoals int
	MatchWeek     int
	Season        string
}

func NewMatch(home, away, season string, week int) Match {
	return Match{
		HomeTeam:  home,
		AwayTeam:  away,
		MatchWeek: week,
		Season:    season,
	}
}

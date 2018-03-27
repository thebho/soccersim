package model

// MatchesDataStore interface
type MatchesDataStore interface {
	GetMatches(string, int) []Match //season, week
}

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
	Played        bool
}

// NewMatch constructor
func NewMatch(home, away, season string, week int) Match {
	return Match{
		HomeTeam:  home,
		AwayTeam:  away,
		MatchWeek: week,
		Season:    season,
	}
}

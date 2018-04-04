package model

// SimWeekRequest is a json request to sim all matches in a season/week
type SimWeekRequest struct {
	SeasonName string
	Week       int
	Action     string
}

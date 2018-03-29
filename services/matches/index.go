package matches

import "soccersim/model"

// MatchService interface for match logic
type MatchService interface {
	GetWeeksMatches(string, int) []model.Match
	SimMatchWeek(string, int)
	ScheduleSeason(string) []model.Match
}

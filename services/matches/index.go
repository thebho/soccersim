package matches

import "github.com/thebho/soccersim/model"

// MatchService interface for match logic
type MatchService interface {
	GetWeeksMatches(string, int) []model.Match
	SimMatchWeek(string, int)
	ScheduleSeason(string) []model.Match
}

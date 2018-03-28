package services

import "SoccerSim/model"

// MatchServices interface for match logic
type MatchServices interface {
	GetWeeksMatches(model.MatchesDataStore, string, int) []model.Match
	SimMatchWeek(MatchWeekDataStore, string, int)
	SimGame(model.Team, model.Team, model.Match)
	ScheduleSeason(SchedulerDataStore, string) []model.Match
}

// TeamService interface for team logic
type TeamService interface {
	GetAllTeams() []model.Team
	GetTeam(string) model.Team
}

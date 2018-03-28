package services

import "SoccerSim/model"

// MatchService interface for match logic
type MatchService interface {
	GetWeeksMatches(MatchDataStore, string, int) []model.Match
	SimMatchWeek(MatchDataStore, string, int)
	SimGame(model.Team, model.Team, model.Match)
	ScheduleSeason(SchedulerDataStore, string) []model.Match
}

// TeamService interface for team logic
type TeamService interface {
	GetAllTeams() []model.Team
	GetTeam(string) model.Team
}

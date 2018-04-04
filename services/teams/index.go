package teams

import "github.com/thebho/soccersim/model"

// TeamService interface for team logic
type TeamService interface {
	GetAllTeams() []model.Team
	GetTeam(string) model.Team
	GetTeamSeason(string, string) model.TeamSeason
	GetTeamSeasons() []model.TeamSeason
}

package teams

import "soccersim/model"

// TeamService interface for team logic
type TeamService interface {
	GetAllTeams() []model.Team
	GetTeam(string) model.Team
}

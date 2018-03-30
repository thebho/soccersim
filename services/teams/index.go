package teams

import "github.com/thebho/soccersim/model"

// TeamService interface for team logic
type TeamService interface {
	GetAllTeams(string) []model.Team
	GetTeam(string) model.Team
}

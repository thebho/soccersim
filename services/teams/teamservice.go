package teams

import "github.com/thebho/soccersim/model"

// TeamDataStore interface
type TeamDataStore interface {
	GetTeam(string) model.Team            //teamABV
	GetTeamsBySeason(string) []model.Team //season
	GetTeams() []model.Team
}

// TeamServiceImp implements TeamServices
type TeamServiceImp struct {
	dataStore TeamDataStore
}

// NewTeamService creates a new TeamService with the provide datastore
func NewTeamService(dataStore TeamDataStore) TeamServiceImp {
	return TeamServiceImp{dataStore: dataStore}
}

// GetAllTeams service.  If season is nil, returns all teams
func (t TeamServiceImp) GetAllTeams(season string) []model.Team {
	if season == "" {
		return t.dataStore.GetTeams()
	}
	return t.dataStore.GetTeamsBySeason(season)
}

// GetTeam service
func (t TeamServiceImp) GetTeam(teamABV string) model.Team {
	return t.dataStore.GetTeam(teamABV)
}

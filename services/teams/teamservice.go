package teams

import "soccersim/model"

// TeamDataStore interface
type TeamDataStore interface {
	GetTeam(string) model.Team //teamABV
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

// GetAllTeams service
func (t TeamServiceImp) GetAllTeams() []model.Team {
	return t.dataStore.GetTeams()
}

// GetTeam service
func (t TeamServiceImp) GetTeam(teamABV string) model.Team {
	return t.dataStore.GetTeam(teamABV)
}

package teams

import (
	"github.com/thebho/soccersim/model"
)

// TeamDataStore interface
type TeamDataStore interface {
	GetTeam(string) model.Team //teamABV
	GetTeams() []model.Team
	GetTeamSeason(string, string) model.TeamSeason
	GetTeamSeasons(string) []model.TeamSeasonJoin
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

// GetTeamSeason returns a TeamSeason (with Team) for the teamABV and season keys
func (t TeamServiceImp) GetTeamSeason(teamABV, season string) model.TeamSeason {
	return t.dataStore.GetTeamSeason(teamABV, season)
}

// GetTeamSeasons all team seasons
func (t TeamServiceImp) GetTeamSeasons(season string) []model.TeamSeasonJoin {
	return t.dataStore.GetTeamSeasons(season)
}

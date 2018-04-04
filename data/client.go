package data

import (
	"github.com/thebho/soccersim/model"
	"github.com/thebho/soccersim/util"

	"github.com/go-pg/pg"
)

// PostgresDS implements datastore interfaces
type PostgresDS struct {
	db *pg.DB
}

// NewPostgresDS constructor
func NewPostgresDS() *PostgresDS {
	return &PostgresDS{db: ConnectToDB()}
}

// GetMatches returns the matches from a season/week
func (p PostgresDS) GetMatches(season string, week int) []model.Match {
	var matches []model.Match
	err := p.db.Model(&matches).
		Where("season = ?", season).
		Where("match_week = ?", week).
		Select()
	util.CheckError(err)
	return matches
}

// GetTeam imp
func (p PostgresDS) GetTeam(teamABV string) model.Team {
	var team model.Team
	err := p.db.Model(&team).
		Where("Abv = ?", teamABV).
		Select()
	util.CheckError(err)
	return team
}

// GetTeamSeason returns a TeamSeason with a Team object
func (p PostgresDS) GetTeamSeason(teamABV, season string) model.TeamSeason {
	var teamSeason model.TeamSeason
	err := p.db.Model(&teamSeason).
		Where("team_id = ?", teamABV).
		Where("season = ?", season).
		First()
	util.CheckError(err)
	// team := p.GetTeam(teamABV)
	// teamSeason.Team = &team
	return teamSeason
}

// GetTeamSeasons returns all team seasons
func (p PostgresDS) GetTeamSeasons(season string) []model.TeamSeasonJoin {
	var teamSeasons []model.TeamSeasonJoin
	query := `SELECT * from team_season INNER JOIN teams ON (team_season.team_id = teams.abv) AND season = ?;`
	_, err := p.db.Model((*model.TeamSeason)(nil)).Query(&teamSeasons, query, season)
	util.CheckError(err)
	return teamSeasons
}

// GetTeams imp
func (p PostgresDS) GetTeams() []model.Team {
	var teams []model.Team
	err := p.db.Model(&teams).Select()
	util.CheckError(err)
	return teams
}

// SaveObject imp
func (p PostgresDS) SaveObject(object interface{}) {
	err := p.db.Insert(object)
	util.CheckError(err)
}

// UpdateObject imp
func (p PostgresDS) UpdateObject(object interface{}) {
	_, err := p.db.Model(object).UpdateNotNull()
	util.CheckError(err)
}

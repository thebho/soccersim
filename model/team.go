package model

// Team data model
type Team struct {
	Abv         string `sql:",pk"`
	Name        string
	TeamSeasons []TeamSeason
}

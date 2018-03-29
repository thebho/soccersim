package util

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"soccersim/model"
)

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// GetTestTeams reads mock team data for testing
func GetTestTeams() []model.Team {
	absPath, err := filepath.Abs("../../util/teams.txt")
	CheckError(err)
	file, err := os.Open(absPath)
	CheckError(err)
	defer file.Close()

	var teams = []model.Team{}
	err = Unmarshal(file, &teams)
	CheckError(err)
	return teams
}

// GetTestTeamsPath reads mock team data for testing
func GetTestTeamsPath(path string) []model.Team {
	absPath, err := filepath.Abs(path)
	CheckError(err)
	file, err := os.Open(absPath)
	CheckError(err)
	defer file.Close()

	var teams = []model.Team{}
	err = Unmarshal(file, &teams)
	CheckError(err)
	return teams
}

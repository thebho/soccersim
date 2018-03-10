package util

import (
	"SoccerSim/model"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// GetTestTeams reads mock team data for testing
func GetTestTeams() []model.Team {
	absPath, _ := filepath.Abs("../util/teams.txt")
	file, err := os.Open(absPath)
	CheckError(err)
	defer file.Close()

	var teams = []model.Team{}
	err = Unmarshal(file, &teams)
	CheckError(err)
	return teams
}

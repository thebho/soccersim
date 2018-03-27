package handlers

import (
	"SoccerSim/model"
	"SoccerSim/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// GetTeams returns all teams from the database
func (s SoccerSim) GetTeams(w http.ResponseWriter, r *http.Request) {
	//TODO: Add logger
	fmt.Println("Getting all teams")
	teams := s.db.GetTeams()
	// writeToFile(teams)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(teams)
}

// for testing
func writeToFile(teams []model.Team) {
	f, err := os.Create("teams.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// w := bufio.NewWriter(f)
	// defer w.Flush()

	r, err := json.Marshal(teams)
	util.CheckError(err)
	_, err = io.Copy(f, bytes.NewReader(r))
	util.CheckError(err)
}

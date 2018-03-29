package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/thebho/soccersim/model"
	"github.com/thebho/soccersim/util"
)

// ScheduleSeason returns all teams from the sdatabase
func (s SoccerSim) ScheduleSeason(w http.ResponseWriter, r *http.Request) {
	//TODO: Add logger
	fmt.Println("Scheduling Season")
	seasonName, err := util.ParseURL("season_name", r.URL.RequestURI())
	util.CheckError(err)
	matches := s.MatchService.ScheduleSeason(seasonName)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(matches)
}

// GetWeeksMatches returns matches for the season/week
func (s SoccerSim) GetWeeksMatches(w http.ResponseWriter, r *http.Request) {
	seasonName, err := util.ParseURL("season_name", r.URL.RequestURI())
	util.CheckError(err)
	week, err := util.ParseURL("week", r.URL.RequestURI())
	util.CheckError(err)
	weekInt, err := strconv.Atoi(week)
	util.CheckError(err)
	fmt.Printf("Getting matches for: %s Week: %d\n", seasonName, weekInt)
	matches := s.MatchService.GetWeeksMatches(seasonName, weekInt)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(matches)
}

// SimWeeksMatches takes a season and week and sims those matches
func (s SoccerSim) SimWeeksMatches(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var simWeekRequest model.SimWeekRequest
	err := decoder.Decode(&simWeekRequest)
	util.CheckError(err)
	defer r.Body.Close()
	if simWeekRequest.Action == "simWeek" {
		s.MatchService.SimMatchWeek(simWeekRequest.SeasonName, simWeekRequest.Week)
	} else {
		response := fmt.Sprintf("Action: %s unknown\n", simWeekRequest.Action)
		http.Error(w, response, http.StatusNotAcceptable)
	}
}

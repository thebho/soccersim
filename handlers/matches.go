package handlers

import (
	"SoccerSim/model"
	"SoccerSim/services"
	"SoccerSim/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ScheduleSeason returns all teams from the database
func (s SoccerSim) ScheduleSeason(w http.ResponseWriter, r *http.Request) {
	//TODO: Add logger
	fmt.Println("Scheduling Season")
	seasonName, err := util.ParseURL("season_name", r.URL.RequestURI())
	util.CheckError(err)
	matches := services.ScheduleSeason(s.db, seasonName)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(matches)
}

// func (s SoccerSim) getWeeksMatches(seasonName string, week int) []model.Match {
// 	var matches []model.Match
// 	err := s.db.Model(&matches).
// 		Where("season = ?", seasonName).
// 		Where("match_week = ?", week).
// 		Select()
// 	util.CheckError(err)
// 	return matches
// }

func (s SoccerSim) GetWeeksMatches(w http.ResponseWriter, r *http.Request) {
	seasonName, err := util.ParseURL("season_name", r.URL.RequestURI())
	util.CheckError(err)
	week, err := util.ParseURL("week", r.URL.RequestURI())
	util.CheckError(err)
	weekInt, err := strconv.Atoi(week)
	util.CheckError(err)
	fmt.Printf("Getting matches for: %s Week: %d\n", seasonName, weekInt)
	matches := s.db.GetMatches(seasonName, weekInt)
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
		services.SimMatchWeek(s.db, simWeekRequest.SeasonName, simWeekRequest.Week)
	} else {
		fmt.Printf("Action: %s unknown\n", simWeekRequest.Action)
	}
}

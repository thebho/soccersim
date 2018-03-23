package controllers

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
func ScheduleSeason(w http.ResponseWriter, r *http.Request) {
	//TODO: Add logger
	fmt.Println("Scheduling Season")
	scheduler := services.SchedulerImp{}
	seasonName, err := util.ParseURL("season_name", r.URL.RequestURI())
	util.CheckError(err)
	matches := scheduler.ScheduleSeason(seasonName)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(matches)
}

func GetWeeksMatches(w http.ResponseWriter, r *http.Request) {
	seasonName, err := util.ParseURL("season_name", r.URL.RequestURI())
	util.CheckError(err)
	week, err := util.ParseURL("week", r.URL.RequestURI())
	util.CheckError(err)
	weekInt, err := strconv.Atoi(week)
	util.CheckError(err)
	fmt.Printf("Getting matches for: %s Week: %d\n", seasonName, weekInt)
	var matches []model.Match
	err = controller.getDB().Model(&matches).
		Where("season = ?", seasonName).
		Where("match_week = ?", weekInt).
		Select()
	util.CheckError(err)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(matches)
}

// SimWeeksMatches takes a season and week and sims those matches
func SimWeeksMatches(w http.ResponseWriter, r *http.Request) {
	seasonName, err := util.ParseURL("season_name", r.URL.RequestURI())
	util.CheckError(err)
	fmt.Println(seasonName)
}

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

func getWeeksMatches(seasonName string, week int) []model.Match {
	var matches []model.Match
	err := controller.getDB().Model(&matches).
		Where("season = ?", seasonName).
		Where("match_week = ?", week).
		Select()
	util.CheckError(err)
	return matches
}

func GetWeeksMatches(w http.ResponseWriter, r *http.Request) {
	seasonName, err := util.ParseURL("season_name", r.URL.RequestURI())
	util.CheckError(err)
	week, err := util.ParseURL("week", r.URL.RequestURI())
	util.CheckError(err)
	weekInt, err := strconv.Atoi(week)
	util.CheckError(err)
	fmt.Printf("Getting matches for: %s Week: %d\n", seasonName, weekInt)
	matches := getWeeksMatches(seasonName, weekInt)
	setReturnDefaults(w)
	json.NewEncoder(w).Encode(matches)
}

// SimWeeksMatches takes a season and week and sims those matches
func SimWeeksMatches(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var simWeekRequest model.SimWeekRequest
	err := decoder.Decode(&simWeekRequest)
	util.CheckError(err)
	defer r.Body.Close()
	if simWeekRequest.Action == "simWeek" {
		matches := getWeeksMatches(simWeekRequest.SeasonName, simWeekRequest.Week)
		for _, match := range matches {
			var homeTeam model.Team
			err = controller.getDB().Model(&homeTeam).
				Where("Abv = ?", match.HomeTeam).
				Select()
			util.CheckError(err)
			var awayTeam model.Team
			err = controller.getDB().Model(&awayTeam).
				Where("Abv = ?", match.AwayTeam).
				Select()
			util.CheckError(err)
			/* TODO: Rework sim game to...
			pull in match and database
			confirm match hasn't been played already
			find teams from the DB
			sim game
			update match in DB
			update Teams in DB
			*/
			services.SimGame(&homeTeam, &awayTeam)
			fmt.Printf("Home Team: %v, Away Team: %v\n", homeTeam, awayTeam)

		}
	} else {
		fmt.Printf("Action: %s unknown\n", simWeekRequest.Action)
	}
}

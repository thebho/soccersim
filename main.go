package main

import (
	"SoccerSim/model"
	"SoccerSim/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var teams = []*model.Team{}

func main() {
	fmt.Println("Running Soccer Sim")
	sim()
}

func sim() {
	teams = append(teams, model.NewTeam("Arsenal"))
	teams = append(teams, model.NewTeam("Totthenham Hotspur"))
	teams = append(teams, model.NewTeam("Manchester United"))
	for _, homeTeam := range teams {
		for _, awayTeam := range teams {
			if homeTeam == awayTeam {
				continue
			}
			services.SimGame(homeTeam, awayTeam)
		}
	}
	for _, team := range teams {
		fmt.Println(team)
	}
}

func writeFile() {
	homeTeam := model.NewTeam("Arsenal")
	homeTeamMarshall, err := json.Marshal(homeTeam)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("/tmp/dat1", homeTeamMarshall, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

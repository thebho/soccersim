package main

import (
	"SoccerSim/model"
	"SoccerSim/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

var teams = []*model.Team{}

func main() {
	fmt.Println("Running Soccer Sim")
	sim()
}

func sim() {
	teams = append(teams, model.NewTeam("Arsenal FC"))
	teams = append(teams, model.NewTeam("Tottenham Hotspur"))
	teams = append(teams, model.NewTeam("Manchester United"))
	teams = append(teams, model.NewTeam("Manchester City"))
	teams = append(teams, model.NewTeam("Leicester City"))
	teams = append(teams, model.NewTeam("Everton"))
	teams = append(teams, model.NewTeam("Chelsea"))
	teams = append(teams, model.NewTeam("Stoke City"))
	teams = append(teams, model.NewTeam("Liverpool"))
	teams = append(teams, model.NewTeam("AFC Bournemouth"))
	teams = append(teams, model.NewTeam("Burnley"))
	teams = append(teams, model.NewTeam("Watford"))
	teams = append(teams, model.NewTeam("West Ham United"))
	teams = append(teams, model.NewTeam("Brighton & Hove Albion"))
	teams = append(teams, model.NewTeam("Huddersfield Town"))
	teams = append(teams, model.NewTeam("Newcastle United"))
	teams = append(teams, model.NewTeam("Southampton"))
	teams = append(teams, model.NewTeam("Swansea City"))
	teams = append(teams, model.NewTeam("West Bromwich Albion"))
	teams = append(teams, model.NewTeam("Crystal Palace"))

	for _, homeTeam := range teams {
		for _, awayTeam := range teams {
			if homeTeam == awayTeam {
				continue
			}
			services.SimGame(homeTeam, awayTeam)
		}
	}

	sort.Sort(model.PointsSorter(teams))

	for rank, team := range teams {
		fmt.Printf("%2d   %-25s %2d %2d %2d -%3d\n", rank+1, team.Name(), team.GetWins(), team.GetDraws(), team.GetLoses(), team.TotalPoints())
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

package main

import (
	"SoccerSim/model"
	"SoccerSim/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Running Soccer Sim")

	sim()

}

func sim() {
	homeTeam := model.NewTeam("Arsenal")
	awayTeam := model.NewTeam("Sp*rs")
	for i := 0; i < 10; i++ {
		services.SimGame(homeTeam, awayTeam)
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

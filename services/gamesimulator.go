package services

import (
	"SoccerSim/model"
	"fmt"

	rand "github.com/thebho/random-tools"
)

func SimGame(homeTeam, awayTeam model.Team) {
	num := rand.RandIndex(3)
	switch num {
	case 0:
		fmt.Printf("%s Wins!\n", homeTeam.Name())
	case 1:
		fmt.Println("Draw!")
	case 2:
		fmt.Printf("%s Wins!\n", awayTeam.Name())
	}
}

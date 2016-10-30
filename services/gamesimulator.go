package services

import (
	"fmt"

	rand "github.com/thebho/random-tools"
)

func SimGame() {
	num := rand.RandInt(10)
	fmt.Println(num)
}

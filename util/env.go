package util

import "github.com/subosito/gotenv"

func LoadEnv() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}
}

package util

import "github.com/subosito/gotenv"

// LoadEnv pulls in .env variables
func LoadEnv() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}
}

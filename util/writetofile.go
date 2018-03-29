package util

import (
	"soccersim/model"
	"bytes"
	"encoding/json"
	"io"
	"os"
)

// WriteToFile for testing
func WriteToFile(teams []model.Team) {
	f, err := os.Create("teams.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// w := bufio.NewWriter(f)
	// defer w.Flush()

	r, err := json.Marshal(teams)
	CheckError(err)
	_, err = io.Copy(f, bytes.NewReader(r))
	CheckError(err)
}

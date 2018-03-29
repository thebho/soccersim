package util

import "testing"

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func TestError(t *testing.T, e error) {
	if e != nil {
		t.Fatal(e)
	}
}

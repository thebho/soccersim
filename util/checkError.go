package util

import "testing"

// CheckError checks an error object for nil and panics if it isn't
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// TestError checks and error for nil and t.Fatals it if it isn't
func TestError(t *testing.T, e error) {
	if e != nil {
		t.Fatal(e)
	}
}

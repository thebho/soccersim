package services

import (
	"testing"
)

func TestSimMatchWeekNoPanic(t *testing.T) {
	SimMatchWeek(mockMatchWeekDataStore, "Test", 1)
}

package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeeksMatches(t *testing.T) {
	matches := GetWeeksMatches(mockMatchDataStore, "MOCK", 1)
	assert.Equal(t, matchA.HomeTeam, matches[0].HomeTeam)
}

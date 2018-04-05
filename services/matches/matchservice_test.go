package matches

import (
	"testing"

	"github.com/thebho/soccersim/model"

	"github.com/stretchr/testify/assert"
)

func TestGetWeeksMatches(t *testing.T) {
	mockMDS := new(MockMatchDataStore)
	mockMDS.On("GetMatches", "MOCK", 1).Return([]model.Match{matchA, matchB})
	matchService := NewMatchService(mockMDS)

	matches := matchService.GetWeeksMatches("MOCK", 1)
	assert.Equal(t, matchA.HomeTeam, matches[0].HomeTeam)
}

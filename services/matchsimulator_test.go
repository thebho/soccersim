package services

import (
	"SoccerSim/model"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestSimMatchWeekNoPanic(t *testing.T) {
	mockMDS := new(MockMatchDataStore)
	mockMDS.On("GetMatches", "MOCK", 1).Return([]model.Match{matchA, matchB})
	mockMDS.On("GetTeam", "ARS").Return(teamA)
	mockMDS.On("GetTeam", mock.Anything).Return(teamB)
	mockMDS.On("UpdateObject", mock.Anything)
	matchService := NewMatchService(mockMDS)
	matchService.SimMatchWeek("MOCK", 1)
}

func TestSimMatchPlayed(t *testing.T) {
	mockMDS := new(MockMatchDataStore)
	mockMDS.On("GetMatches", "MOCK", 1).Return([]model.Match{matchA, matchB})
	matchService := NewMatchService(mockMDS)

	// This will blow up if false because there is no datastore
	match := model.Match{Played: true}
	matchService = NewMatchService(mockMDS)
	matchService.simMatch(match)
	mockMDS.AssertNotCalled(t, "GetTeam", mock.Anything)
}

func TestSimMatch(t *testing.T) {

	mockMDS := new(MockMatchDataStore)
	match := model.Match{HomeTeam: "TeamA", AwayTeam: "TeamB", Played: false}
	mockMDS.On("GetTeam", "TeamA").Return(teamA)
	mockMDS.On("GetTeam", "TeamB").Return(teamB)
	mockMDS.On("UpdateObject", mock.Anything)
	matchService := NewMatchService(mockMDS)

	matchService.simMatch(match)
	mockMDS.AssertCalled(t, "GetTeam", "TeamA")
	mockMDS.AssertCalled(t, "GetTeam", "TeamB")
	mockMDS.AssertNumberOfCalls(t, "UpdateObject", 3)
}

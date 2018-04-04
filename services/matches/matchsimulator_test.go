package matches

import (
	"testing"

	"github.com/thebho/soccersim/model"

	"github.com/stretchr/testify/mock"
)

func TestSimMatchWeekNoPanic(t *testing.T) {
	mockMDS := new(MockMatchDataStore)
	mockMDS.On("GetMatches", "MOCK", 1).Return([]model.Match{matchA, matchB})
	mockMDS.On("GetTeamSeason", "ARS", "MOCK").Return(teamA)
	mockMDS.On("GetTeamSeason", mock.Anything, "MOCK").Return(teamB)
	mockMDS.On("UpdateObject", mock.Anything)
	mockSimulator := new(MockMatchSimulator)
	mockSimulator.On("Sim", mock.Anything, mock.Anything, mock.Anything)
	matchService := NewMatchService(mockMDS)
	matchService.matchSimulator = mockSimulator
	matchService.SimMatchWeek("MOCK", 1)
}

func TestSimMatchPlayed(t *testing.T) {
	mockMDS := new(MockMatchDataStore)
	mockMDS.On("GetMatches", "MOCK", 1).Return([]model.Match{matchA, matchB})
	matchService := NewMatchService(mockMDS)

	// This will blow up if false because there is no datastore
	match := model.Match{Played: true}
	matchService = NewMatchService(mockMDS)
	matchService.simMatch(match, "MOCK")
	mockMDS.AssertNotCalled(t, "GetTeamSeason", mock.Anything, mock.Anything)
}

func TestSimMatch(t *testing.T) {

	mockMDS := new(MockMatchDataStore)
	match := model.Match{HomeTeam: "TeamA", AwayTeam: "TeamB", Played: false}
	mockMDS.On("GetTeamSeason", "TeamA", "MOCK").Return(teamA)
	mockMDS.On("GetTeamSeason", "TeamB", "MOCK").Return(teamB)
	mockMDS.On("UpdateObject", mock.Anything)
	mockSimulator := new(MockMatchSimulator)
	mockSimulator.On("Sim", mock.Anything, mock.Anything, mock.Anything)
	matchService := NewMatchService(mockMDS)
	matchService.matchSimulator = mockSimulator

	matchService.simMatch(match, "MOCK")
	mockMDS.AssertCalled(t, "GetTeamSeason", "TeamA", "MOCK")
	mockMDS.AssertCalled(t, "GetTeamSeason", "TeamB", "MOCK")
	mockMDS.AssertNumberOfCalls(t, "UpdateObject", 3)
}

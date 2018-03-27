package services

import (
	"SoccerSim/model"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestSimMatchWeekNoPanic(t *testing.T) {
	SimMatchWeek(mockMatchWeekDataStore, "Test", 1)
}

func TestSimMatchPlayed(t *testing.T) {
	testObj := new(MyMockedObject)

	// This will blow up if false because there is no datastore
	match := model.Match{Played: true}
	simulator2 := simulator{testObj}
	simulator2.simMatch(match)
	testObj.AssertNotCalled(t, "GetTeam", mock.Anything)
}

func TestSimMatch(t *testing.T) {
	testObj := new(MyMockedObject)
	match := model.Match{HomeTeam: "TeamA", AwayTeam: "TeamB", Played: false}
	testObj.On("GetTeam", "TeamA").Return(teamA)
	testObj.On("GetTeam", "TeamB").Return(teamB)
	testObj.On("UpdateObject", mock.Anything)

	simulator2 := simulator{testObj}
	simulator2.simMatch(match)
	testObj.AssertCalled(t, "GetTeam", "TeamA")
	testObj.AssertCalled(t, "GetTeam", "TeamB")
	testObj.AssertNumberOfCalls(t, "UpdateObject", 3)
}

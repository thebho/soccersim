package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thebho/soccersim/handlers"
	"github.com/thebho/soccersim/util"
)

func TestGetTeamsNoSeason(t *testing.T) {
	req, err := http.NewRequest("GET", "/teams", nil)
	util.TestError(t, err)
	mockTS := new(MockTeamService)
	mockTS.On("GetAllTeams", "")
	mockSoccerSim := handlers.NewSoccerSim()
	mockSoccerSim.TeamService = mockTS
	mockSoccerSim.MatchService = nil
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockSoccerSim.GetTeams)
	handler.ServeHTTP(rr, req)
}

func TestGetTeamsSeason(t *testing.T) {
	req, err := http.NewRequest("GET", "/teams?season=mock", nil)
	util.TestError(t, err)
	mockTS := new(MockTeamService)
	mockTS.On("GetAllTeams", "mock")
	mockSoccerSim := handlers.NewSoccerSim()
	mockSoccerSim.TeamService = mockTS
	mockSoccerSim.MatchService = nil
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockSoccerSim.GetTeams)
	handler.ServeHTTP(rr, req)
}

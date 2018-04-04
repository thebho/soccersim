package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/thebho/soccersim/handlers"
	"github.com/thebho/soccersim/util"
)

func TestGetTeamSeason(t *testing.T) {
	season := "MOCK_SEASON"
	team := "MKA"
	req, err := http.NewRequest("GET", "/seasons/MOCK_SEASON/teams/MKA", nil)
	req = mux.SetURLVars(req, map[string]string{"seasonKey": season, "teamAbv": team})
	util.TestError(t, err)
	mockTS := new(MockTeamService)
	mockTS.On("GetTeamSeason", team, season)
	mockSoccerSim := handlers.NewSoccerSim()
	mockSoccerSim.TeamService = mockTS
	mockSoccerSim.MatchService = nil
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockSoccerSim.GetTeamSeason)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestGetSeasonTeams(t *testing.T) {
	season := "MOCK_SEASON"
	req, err := http.NewRequest("GET", "/seasons/MOCK_SEASON/teams/", nil)
	req = mux.SetURLVars(req, map[string]string{"seasonKey": season})
	util.TestError(t, err)
	mockTS := new(MockTeamService)
	mockTS.On("GetTeamSeasons", season)
	mockSoccerSim := handlers.NewSoccerSim()
	mockSoccerSim.TeamService = mockTS
	mockSoccerSim.MatchService = nil
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockSoccerSim.GetSeasonTeams)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/thebho/soccersim/handlers"
	"github.com/thebho/soccersim/model"
	"github.com/thebho/soccersim/util"

	"github.com/stretchr/testify/assert"
)

func TestScheduleSeason(t *testing.T) {
	req, err := http.NewRequest("POST", "/season?season_name=Test", nil)
	util.TestError(t, err)
	mockMS := new(MockMatchService)
	mockMS.On("ScheduleSeason", "Test")
	mockSoccerSim := handlers.NewSoccerSim()
	mockSoccerSim.MatchService = mockMS
	mockSoccerSim.TeamService = nil
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockSoccerSim.ScheduleSeason)
	handler.ServeHTTP(rr, req)
	assert.NotNil(t, rr.Body)
	assert.Equal(t, 200, rr.Code)
}

func TestGetWeeksMatches(t *testing.T) {
	req, err := http.NewRequest("GET", "/matches?season_name=Test&week=1", nil)
	util.TestError(t, err)
	mockMS := new(MockMatchService)
	mockMS.On("GetWeeksMatches", "Test", 1)
	mockSoccerSim := handlers.NewSoccerSim()
	mockSoccerSim.MatchService = mockMS
	mockSoccerSim.TeamService = nil
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockSoccerSim.GetWeeksMatches)
	handler.ServeHTTP(rr, req)
	assert.NotNil(t, rr.Body)
	assert.Equal(t, 200, rr.Code)
}

func TestSimWeeksMatches(t *testing.T) {
	mockSimRequest := model.SimWeekRequest{
		Action:     "simWeek",
		SeasonName: "Test",
		Week:       1,
	}
	mockRequestEncoded, err := json.Marshal(mockSimRequest)
	util.CheckError(err)
	req, err := http.NewRequest("POST", "/matches", strings.NewReader(string(mockRequestEncoded)))
	util.TestError(t, err)
	mockMS := new(MockMatchService)
	mockMS.On("SimMatchWeek", "Test", 1)
	mockSoccerSim := handlers.NewSoccerSim()
	mockSoccerSim.MatchService = mockMS
	mockSoccerSim.TeamService = nil
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockSoccerSim.SimWeeksMatches)
	handler.ServeHTTP(rr, req)
	assert.NotNil(t, rr.Body)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestBadAction(t *testing.T) {
	mockSimRequest := model.SimWeekRequest{
		Action:     "bad",
		SeasonName: "Test",
		Week:       1,
	}
	mockRequestEncoded, err := json.Marshal(mockSimRequest)
	util.CheckError(err)
	req, err := http.NewRequest("POST", "/matches", strings.NewReader(string(mockRequestEncoded)))
	util.TestError(t, err)
	mockMS := new(MockMatchService)
	mockMS.On("SimMatchWeek", "Test", 1)
	mockSoccerSim := handlers.NewSoccerSim()
	mockSoccerSim.MatchService = mockMS
	mockSoccerSim.TeamService = nil
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockSoccerSim.SimWeeksMatches)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotAcceptable, rr.Code)
}

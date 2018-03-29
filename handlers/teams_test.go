package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"soccersim/handlers"
	"soccersim/util"
	"testing"
)

// GetTeams returns all teams from the database
func TestGetTeams(t *testing.T) {
	req, err := http.NewRequest("GET", "/teams", nil)
	util.TestError(t, err)
	mockTS := new(MockTeamService)
	mockTS.On("GetAllTeams")
	mockSoccerSim := handlers.NewSoccerSim()
	mockSoccerSim.TeamService = mockTS
	mockSoccerSim.MatchService = nil
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockSoccerSim.GetTeams)
	handler.ServeHTTP(rr, req)
	// teams := mockSoccerSim.GetAllTeams()
	// // util.WriteToFile(teams)
	// setReturnDefaults(w)
	// json.NewEncoder(w).Encode(teams)
}

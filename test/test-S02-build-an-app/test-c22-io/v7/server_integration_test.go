package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	inputoutput "learn.go/S02-build-an-app/c22-io/v7"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := CreateTempFile(t, "")
	defer cleanDatabase()
	store := inputoutput.NewFileSystemPlayerStore(database)
	server := inputoutput.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), inputoutput.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), inputoutput.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), inputoutput.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, inputoutput.NewGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, inputoutput.NewLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []inputoutput.Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
}

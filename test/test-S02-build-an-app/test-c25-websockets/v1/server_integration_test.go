package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	poker "learn.go/S02-build-an-app/c25-websockets/v1"

)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := CreateTempFile(t, `[]`)
	defer cleanDatabase()
	store, err := poker.NewFileSystemPlayerStore(database)

	assertNoError(t, err)

	server := mustMakePlayerServer(t, store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get League", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
}

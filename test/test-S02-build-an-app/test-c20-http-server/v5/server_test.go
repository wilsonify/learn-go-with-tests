package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	httpserver "learn.go/S02-build-an-app/c20-http-server/v5"
)

func TestGETPlayers(t *testing.T) {
	store := httpserver.StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
	}
	server := &httpserver.PlayerServer{&store}

	tests := []struct {
		name               string
		player             string
		expectedHTTPStatus int
		expectedScore      string
	}{
		{
			name:               "Returns Pepper's score",
			player:             "Pepper",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "20",
		},
		{
			name:               "Returns Floyd's score",
			player:             "Floyd",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "10",
		},
		{
			name:               "Returns 404 on missing players",
			player:             "Apollo",
			expectedHTTPStatus: http.StatusNotFound,
			expectedScore:      "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httpserver.NewGetScoreRequest(tt.player)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			assertStatus(t, response.Code, tt.expectedHTTPStatus)
			assertResponseBody(t, response.Body.String(), tt.expectedScore)
		})
	}
}

func TestStoreWins(t *testing.T) {
	store := httpserver.StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &httpserver.PlayerServer{&store}

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"

		request := httpserver.NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.WinningCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.WinningCalls), 1)
		}

		if store.WinningCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.WinningCalls[0], player)
		}
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

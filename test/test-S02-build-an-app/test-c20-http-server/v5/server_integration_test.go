package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	httpserver "learn.go/S02-build-an-app/c20-http-server/v5"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := httpserver.NewInMemoryPlayerStore()
	server := httpserver.PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, httpserver.NewGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	marshal "learn.go/S02-build-an-app/c21-json/v4"

)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := marshal.NewInMemoryPlayerStore()
	server := marshal.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), marshal.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), marshal.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), marshal.NewPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, marshal.NewGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

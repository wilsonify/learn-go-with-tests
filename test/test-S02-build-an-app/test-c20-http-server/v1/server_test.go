package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	httpserver "learn.go/S02-build-an-app/c20-http-server/v1"
)

func TestGETPlayers(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	httpserver.PlayerServer(response, request)

	t.Run("returns Pepper's score", func(t *testing.T) {
		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}

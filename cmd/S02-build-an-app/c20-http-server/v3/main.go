package httpserver

import (
	"log"
	"net/http"

	httpserver "learn.go/S02-build-an-app/c20-http-server/v3"
)

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct{}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &httpserver.PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}

package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		p.processWin(w)
	case http.MethodGet:
		p.showScore(w, r)
	}

}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct{}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func NewGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func Mainly() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}


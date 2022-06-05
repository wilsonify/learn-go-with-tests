package main

import (
	"strings"
	"testing"

	inputoutput "learn.go/S02-build-an-app/c22-io/v1"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		store := inputoutput.FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []inputoutput.Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}

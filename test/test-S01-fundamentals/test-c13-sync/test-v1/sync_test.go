package v1

import (
	"testing"

	syncronize "learn.go/S01-fundamentals/c13-sync/v1"
)

func TestCounter(t *testing.T) {

	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := syncronize.Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
}

func assertCounter(t testing.TB, got syncronize.Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

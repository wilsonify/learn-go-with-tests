package poker

import (
	"testing"
	poker "learn.go/S02-build-an-app/c23-command-line/v2"

)

func TestCLI(t *testing.T) {
	playerStore := &StubPlayerStore{}

	cli := &poker.CLI{playerStore}
	cli.PlayPoker()

	if len(playerStore.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}
}

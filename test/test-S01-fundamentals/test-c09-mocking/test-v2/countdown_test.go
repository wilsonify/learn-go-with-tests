package main

import (
	"bytes"
	"testing"

	mocking "learn.go/S01-fundamentals/c09-mocking/v2"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	mocking.Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

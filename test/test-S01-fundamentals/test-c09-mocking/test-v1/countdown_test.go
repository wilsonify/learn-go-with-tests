package main

import (
	"bytes"
	"testing"

	mocking "learn.go/S01-fundamentals/c09-mocking/v1"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	mocking.Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

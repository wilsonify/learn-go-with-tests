package main

import (
	"../../../../src/S01-fundamentals/c01-hello-world/v2/hello"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

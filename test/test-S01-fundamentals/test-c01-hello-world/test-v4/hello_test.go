package main

import (
	hi "S01-fundamentals/c01-hello-world/v4"
	"testing"
)

func TestHello(t *testing.T) {
	got := hi.Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

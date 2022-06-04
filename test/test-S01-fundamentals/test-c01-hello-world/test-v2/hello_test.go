package main

import (
	hi "S01-fundamentals/c01-hello-world/v2"
	"testing"
)

func TestHello(t *testing.T) {

	got := hi.Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

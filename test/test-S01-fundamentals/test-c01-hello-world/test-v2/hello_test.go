package main

import (
	"testing"

	hi "learn.go/S01-fundamentals/c01-hello-world/v1"
)

func TestHello(t *testing.T) {

	got := hi.Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

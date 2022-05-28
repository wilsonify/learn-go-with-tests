package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHere(t *testing.T) {
	got := here()
	want := "/home/thom/repos/learn-go-with-tests/src/S01-fundamentals/c01-hello-world/v2/"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

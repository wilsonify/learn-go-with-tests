package main

import (
	"testing"

	hello "learn.go/S01-fundamentals/c01-hello-world/v1"
)

func TestHello(t *testing.T) {

	got := hello.Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

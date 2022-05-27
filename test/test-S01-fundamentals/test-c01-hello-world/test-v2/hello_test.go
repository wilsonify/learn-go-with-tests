package main

import (
	Hello "S01-fundamentals/test-c01-hello-world/test-v2/hello/v2"
	"testing"
)

func TestHello(t *testing.T) {

	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

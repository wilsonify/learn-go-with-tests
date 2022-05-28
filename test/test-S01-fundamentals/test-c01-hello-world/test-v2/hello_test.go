package main

import (
	v2 "S01-fundamentals/c01-hello-world/v2/hello.go"
	"testing"
)

func TestHello(t *testing.T) {

	got := v2.Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

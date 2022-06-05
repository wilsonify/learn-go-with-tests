package main

import (
	"bytes"
	"testing"

	injection "learn.go/S01-fundamentals/c08-dependency-injection/v2"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	injection.Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

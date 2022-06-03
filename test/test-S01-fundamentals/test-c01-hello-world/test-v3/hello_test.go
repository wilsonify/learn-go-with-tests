package main

import (
	hi "github.com/wilsonify/learn-go-with-tests/S01-fundamentals/c01-hello-world/v3"
	"testing"
)

func TestHello(t *testing.T) {
	result := hi.Hello("Chris")
	expected := "Hello, Chris"

	if result != expected {
		t.Errorf("got %q want %q", result, expected)
	}
}

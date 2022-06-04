package main

import (
	hello "learn.go/S01-fundamentals/c01-hello-world/v7"
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := hello.Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := hello.Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(got, want)
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		got := hello.Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := hello.Hello("Lauren", "French")
		want := "Bonjour, Lauren"
		assertCorrectMessage(got, want)
	})

}

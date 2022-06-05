package main

import (
	"testing"

	structs "learn.go/S01-fundamentals/c05-structs/v1"
)

func TestPerimeter(t *testing.T) {
	got := structs.Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

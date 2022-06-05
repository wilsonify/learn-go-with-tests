package main

import (
	"testing"

	arrays "learn.go/S01-fundamentals/c04-arrays/v2"
)

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := arrays.Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

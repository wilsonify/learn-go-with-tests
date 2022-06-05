package integers_test

import (
	"testing"

	integers "learn.go/S01-fundamentals/c02-integers/v1"
)

func TestAdder(t *testing.T) {
	sum := integers.Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

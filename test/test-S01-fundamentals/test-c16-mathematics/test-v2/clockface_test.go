package clockface_test

import (
	"math"
	"testing"
	"time"
	clockfacemath "learn.go/S01-fundamentals/c16-mathematics/v2"
)

func TestSecondsInRadians(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := clockfacemath.SecondsInRadians(thirtySeconds)

	if want != got {
		t.Fatalf("Wanted %v radians, but got %v", want, got)
	}
}

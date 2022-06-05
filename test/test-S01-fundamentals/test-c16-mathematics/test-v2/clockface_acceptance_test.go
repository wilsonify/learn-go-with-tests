package clockface_test

import (
	"testing"
	"time"

	clockfacemath "learn.go/S01-fundamentals/c16-mathematics/v2"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockfacemath.Point{X: 150, Y: 150 - 90}
	got := clockfacemath.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

// 	want := clockfacemath.Point{X: 150, Y: 150 + 90}
// 	got := clockfacemath.SecondHand(tm)

// 	if got != want {
// 		t.Errorf("Got %v, wanted %v", got, want)
// 	}
// }

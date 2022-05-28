package clockface_test

import (
	"github.com/wilsonify/learning-go-with-tests/S01-fundamentals/c16-mathematics/v1"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := v1.Point{X: 150, Y: 150 - 90}
	got := v1.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

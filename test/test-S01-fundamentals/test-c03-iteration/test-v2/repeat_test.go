package iteration

import (
	"testing"

	iteration "learn.go/S01-fundamentals/c03-iteration/v2"
)

func TestRepeat(t *testing.T) {
	repeated := iteration.Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

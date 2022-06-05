package iteration

import (
	"testing"

	iteration "learn.go/S01-fundamentals/c03-iteration/vx"
)

func TestRepeat(t *testing.T) {
	repeated := iteration.Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a")
	}
}

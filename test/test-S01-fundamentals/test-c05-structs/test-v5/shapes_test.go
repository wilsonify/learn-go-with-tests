package main

import (
	"testing"

	structs "learn.go/S01-fundamentals/c05-structs/v5"
)

func TestPerimeter(t *testing.T) {
	rectangle := structs.Rectangle{10.0, 10.0}
	got := structs.Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape structs.Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := structs.Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := structs.Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

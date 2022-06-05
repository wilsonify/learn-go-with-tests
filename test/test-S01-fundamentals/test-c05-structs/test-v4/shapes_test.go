package main

import (
	"testing"

	structs "learn.go/S01-fundamentals/c05-structs/v4"
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

	t.Run("rectangles", func(t *testing.T) {
		rectangle := structs.Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := structs.Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}

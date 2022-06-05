package main

import (
	"testing"

	structs "learn.go/S01-fundamentals/c05-structs/v6"
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

	areaTests := []struct {
		shape structs.Shape
		want  float64
	}{
		{structs.Rectangle{12, 6}, 72.0},
		{structs.Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}

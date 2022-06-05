package main

import (
	"testing"

	structs "learn.go/S01-fundamentals/c05-structs/v8"
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
		name    string
		shape   structs.Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: structs.Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: structs.Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: structs.Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}

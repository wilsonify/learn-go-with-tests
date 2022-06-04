package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}

func SecondHandPoint(t time.Time) Point {
	return AngleToPoint(SecondsInRadians(t))
}

func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}

func MinuteHandPoint(t time.Time) Point {
	return AngleToPoint(MinutesInRadians(t))
}

func AngleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

const (
	SecondHandLength = 90
	MinuteHandLength = 80
	clockCentreX     = 150
	clockCentreY     = 150
)

//SVGWriter writes an SVG representation of an analogue clock, showing the time t, to the writer w.
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondHand(w, t)
	MinuteHand(w, t)
	io.WriteString(w, svgEnd)
}

func SecondHand(w io.Writer, t time.Time) {
	p := MakeHand(SecondHandPoint(t), SecondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func MinuteHand(w io.Writer, t time.Time) {
	p := MakeHand(MinuteHandPoint(t), MinuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func MakeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	return Point{p.X + clockCentreX, p.Y + clockCentreY}
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

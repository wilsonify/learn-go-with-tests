package racer

import (
	"net/http"
	"time"
)

// Racer compares the response times of a and b, returning the fastest one.
func Racer(a, b string) (winner string) {
	aDuration := MeasureResponseTime(a)
	bDuration := MeasureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func MeasureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

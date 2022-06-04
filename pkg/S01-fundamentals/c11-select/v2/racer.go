package racer

import (
	"net/http"
)

// Racer compares the response times of a and b, returning the fastest one.
func Racer(a, b string) (winner string) {
	select {
	case <-Ping(a):
		return a
	case <-Ping(b):
		return b
	}
}

func Ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

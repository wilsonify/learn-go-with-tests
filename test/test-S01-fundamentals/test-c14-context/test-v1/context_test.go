package context1

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	contextual "learn.go/S01-fundamentals/c14-context/v1"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

// SpyStore allows you to simulate a store and see how its used.
type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

// Fetch returns response after a short delay.
func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

// Cancel will record the call.
func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}

func TestServer(t *testing.T) {
	data := "hello, world"
	svr := contextual.Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}

package errortypes

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	errortypes "learn.go/S03-q-and-a/c27-error-types/v2"
)

func TestDumbGetter(t *testing.T) {

	t.Run("when you don't get a 200 you get a status error", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusTeapot)
		}))
		defer svr.Close()

		_, err := errortypes.DumbGetter(svr.URL)

		if err == nil {
			t.Fatal("expected an error")
		}

		var got errortypes.BadStatusError
		isBadStatusError := errors.As(err, &got)
		want := errortypes.BadStatusError{URL: svr.URL, Status: http.StatusTeapot}

		if !isBadStatusError {
			t.Fatalf("was not a BadStatusError, got %T", err)
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

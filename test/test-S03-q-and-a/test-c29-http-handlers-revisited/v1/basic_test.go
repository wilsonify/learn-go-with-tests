package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	httphandler "learn.go/S03-q-and-a/c29-http-handlers-revisited/v1"
)

func TestTeapotHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	httphandler.Teapot(res, req)

	if res.Code != http.StatusTeapot {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusTeapot)
	}
}

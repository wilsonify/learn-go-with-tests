package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	httphandler "learn.go/s03/q-and-a/c29-http-handlers-revisited/v1"
)

type MockUserService struct {
	RegisterFunc    func(user httphandler.User) (string, error)
	UsersRegistered []httphandler.User
}

func (m *MockUserService) Register(user httphandler.User) (insertedID string, err error) {
	m.UsersRegistered = append(m.UsersRegistered, user)
	return m.RegisterFunc(user)
}

func TestRegisterUser(t *testing.T) {
	t.Run("can register valid users", func(t *testing.T) {
		user := httphandler.User{Name: "CJ"}
		expectedInsertedID := "whatever"

		service := &MockUserService{
			RegisterFunc: func(user httphandler.User) (string, error) {
				return expectedInsertedID, nil
			},
		}
		server := httphandler.NewUserServer(service)

		req := httptest.NewRequest(http.MethodGet, "/", userToJSON(user))
		res := httptest.NewRecorder()

		server.RegisterUser(res, req)

		assertStatus(t, res.Code, http.StatusCreated)

		if res.Body.String() != expectedInsertedID {
			t.Errorf("expected body of %q but got %q", res.Body.String(), expectedInsertedID)
		}

		if len(service.UsersRegistered) != 1 {
			t.Fatalf("expected 1 user added but got %d", len(service.UsersRegistered))
		}

		if !reflect.DeepEqual(service.UsersRegistered[0], user) {
			t.Errorf("the user registered %+v was not what was expected %+v", service.UsersRegistered[0], user)
		}
	})

	t.Run("returns 400 bad request if body is not valid user JSON", func(t *testing.T) {
		server := httphandler.NewUserServer(nil)

		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("trouble will find me"))
		res := httptest.NewRecorder()

		server.RegisterUser(res, req)

		assertStatus(t, res.Code, http.StatusBadRequest)
	})

	t.Run("returns a 500 internal server error if the service fails", func(t *testing.T) {
		user := httphandler.User{Name: "CJ"}

		service := &MockUserService{
			RegisterFunc: func(user httphandler.User) (string, error) {
				return "", errors.New("couldn't add new user")
			},
		}
		server := httphandler.NewUserServer(service)

		req := httptest.NewRequest(http.MethodGet, "/", userToJSON(user))
		res := httptest.NewRecorder()

		server.RegisterUser(res, req)

		assertStatus(t, res.Code, http.StatusInternalServerError)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted http status %d but got %d", got, want)
	}
}

func userToJSON(user httphandler.User) io.Reader {
	stuff, _ := json.Marshal(user)
	return bytes.NewReader(stuff)
}

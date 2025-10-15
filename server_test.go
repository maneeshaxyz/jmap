package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubUserStore struct {
	strings map[string]string
}

func (s *StubUserStore) GetUserString(name string) string {
	userString := s.strings[name]
	return userString
}

func newGetRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func TestGETRequest(t *testing.T) {
	store := StubUserStore{
		map[string]string{
			"Siripala": "Siripala's Mailbox",
			"Piyaseli": "Piyaseli's Mailbox",
		},
	}

	server := &UserServer{&store}

	t.Run("returns a value for Siripala", func(t *testing.T) {
		request := newGetRequest("Siripala")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "Siripala's Mailbox")
	})

	t.Run("returns a value for Piyaseli", func(t *testing.T) {
		request := newGetRequest("Piyaseli")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "Piyaseli's Mailbox")
	})

	t.Run("returns 404 on missing users", func(t *testing.T) {
		request := newGetRequest("PunchiBanda")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})
}

func TestPOSTRequest(t *testing.T) {
	store := StubUserStore{
		map[string]string{},
	}
	server := &UserServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/users/PunchiBanda", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
	})
}

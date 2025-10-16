package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubUserStore struct {
	strings map[string]string
	values  []string
}

func (s *StubUserStore) GetUserMailbox(name string) string {
	score := s.strings[name]
	return score
}

func (s *StubUserStore) ChangeUserValues(name string) {
	s.values = append(s.values, name)
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
		}, nil,
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
		map[string]string{}, nil,
	}
	server := &UserServer{&store}

	t.Run("it records values on POST", func(t *testing.T) {
		user := "PunchiBanda"

		request := newPostRequest(user)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.values) != 1 {
			t.Fatalf("got %d calls to ChangeValues want %d", len(store.values), 1)
		}

		if store.values[0] != user {
			t.Errorf("did not store correct winner got %q want %q", store.values[0], user)
		}
	})
}

func newPostRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/users/%s", name), nil)
	return req
}

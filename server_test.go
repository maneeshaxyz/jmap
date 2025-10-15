package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func TestGETRequest(t *testing.T) {
	t.Run("returns a value for Siripala", func(t *testing.T) {
		request := newGetRequest("Siripala")
		response := httptest.NewRecorder()

		GetReq(response, request)

		assertResponseBody(t, response.Body.String(), "Siripala's mailbox")
	})

	t.Run("returns a value for Piyaseli", func(t *testing.T) {
		request := newGetRequest("Piyaseli")
		response := httptest.NewRecorder()

		GetReq(response, request)

		assertResponseBody(t, response.Body.String(), "Piyaseli's mailbox")
	})
}

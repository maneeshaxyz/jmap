package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETRequest(t *testing.T) {
	t.Run("returns a GET req for Siripala", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users/Siripala", nil)
		response := httptest.NewRecorder()

		GetReq(response, request)

		got := response.Body.String()
		want := "Siripala's mailbox"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Assert helper functions

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

/// Main Tests

func TestPOSTRequest(t *testing.T) {
	t.Run("it sends a POST value and returns 200", func(t *testing.T) {
		server := &JMAPServer{}

		request, _ := http.NewRequest(http.MethodPost, "/jmap", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
	})
}

package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestJMAPServer(t *testing.T) {
	t.Run("it sends a POST value and returns 200", func(t *testing.T) {
		server := &JMAPServer{}

		request, _ := http.NewRequest(http.MethodPost, "/jmap", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("response contains valid JSON", func(t *testing.T) {
		server := &JMAPServer{}
		request, _ := http.NewRequest(http.MethodPost, "/jmap", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Header().Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type application/json")
		}

		got := response.Body.String()
		if !strings.Contains(got, "methodResponses") {
			t.Errorf("response missing methodResponses")
		}
	})
}

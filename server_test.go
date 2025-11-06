package main

import (
	"encoding/json"
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

// Main tests

func TestJMAPServer(t *testing.T) {
	t.Run("returns 202 Accepted for a valid JMAP request", func(t *testing.T) {
		server := &JMAPServer{}
		requestBody := strings.NewReader(`{"using":[], "methodCalls":[]}`)
		request, _ := http.NewRequest(http.MethodPost, "/jmap", requestBody)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)

		if response.Header().Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type application/json, got %q", response.Header().Get("Content-Type"))
		}

		var jmapResp JMAPResponse
		err := json.Unmarshal(response.Body.Bytes(), &jmapResp)
		if err != nil {
			t.Fatalf("Failed to unmarshal response JSON: %v", err)
		}

		if jmapResp.MethodResponses == nil {
			t.Error("MethodResponses should be an empty slice, not nil")
		}

		if jmapResp.SessionState == "" {
			t.Error("sessionState should not be empty")
		}
	})

	t.Run("returns 400 Bad Request for invalid JSON", func(t *testing.T) {
		server := &JMAPServer{}
		requestBody := strings.NewReader("{this is not valid json")
		request, _ := http.NewRequest(http.MethodPost, "/jmap", requestBody)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusBadRequest)

		// TODO: assert the response body is empty or contains a specific JSON error, per RFC 8620.
	})
}

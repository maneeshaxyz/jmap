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
	t.Run("it sends a POST value and returns 200", func(t *testing.T) {
		server := &JMAPServer{}

		request, _ := http.NewRequest(http.MethodPost, "/jmap", strings.NewReader("{}"))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("RESPONSE contains valid JSON", func(t *testing.T) {
		server := &JMAPServer{}
		request, _ := http.NewRequest(http.MethodPost, "/jmap", strings.NewReader("{}"))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Header().Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type application/json")
		}

		var jmapResp JMAPResponse
		err := json.Unmarshal(response.Body.Bytes(), &jmapResp)
		if err != nil {
			t.Fatalf("JSON unmarshal failed")
		}
		if jmapResp.MethodResponses == nil {
			t.Error("error")
		}
		if jmapResp.SessionState == "" {
			t.Error("sessionState should not be empty")
		}
	})

	t.Run("Read and parse incoming JSON", func(t *testing.T) {
		server := &JMAPServer{}
		requestBody := strings.NewReader(`{"using":[], "methodCalls":[]}`)
		request, _ := http.NewRequest(http.MethodPost, "/jmap", requestBody)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		// If server parsed successfully, it returns 202
		assertStatus(t, response.Code, http.StatusAccepted)
	})

	t.Run("return error for invalid JSON", func(t *testing.T) {
		server := &JMAPServer{}
		requestBody := strings.NewReader("{this is not valid json")
		request, _ := http.NewRequest(http.MethodPost, "/jmap", requestBody)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		// parse and return 400
		assertStatus(t, response.Code, http.StatusBadRequest)
	})
}

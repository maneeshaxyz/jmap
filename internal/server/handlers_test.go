package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maneeshaxyz/jmap/internal/jmap"
)

func TestSessionHandler(t *testing.T) {
	app := newTestApplication(t)

	req, err := http.NewRequest("GET", "/.well-known/jmap", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.sessionHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the Content-Type header
	expectedContentType := "application/json"
	if ctype := rr.Header().Get("Content-Type"); ctype != expectedContentType {
		t.Errorf("Content-Type header does not match: got %s want %s",
			ctype, expectedContentType)
	}

	var session jmap.Session
	if err := json.NewDecoder(rr.Body).Decode(&session); err != nil {
		t.Fatalf("could not decode response body: %s", err)
	}

	if session.Username != jmap.DefaultUsername {
		t.Errorf("username is incorrect: got %s want %s", session.Username, jmap.DefaultUsername)
	}

	if session.APIURL != jmap.APIBaseURL {
		t.Errorf("apiUrl is incorrect: got %s want %s", session.APIURL, jmap.APIBaseURL)
	}

	// Check a few core capability values
	coreCap, ok := session.Capabilities["urn:ietf:params:jmap:core"].(map[string]any)
	if !ok {
		t.Fatal("core capability not found or is not a map")
	}

	if val := coreCap["maxSizeUpload"]; uint64(val.(float64)) != jmap.MaxSizeUpload {
		t.Errorf("maxSizeUpload is incorrect: got %v want %v", val, jmap.MaxSizeUpload)
	}

	if val := coreCap["maxCallsInRequest"]; uint64(val.(float64)) != jmap.MaxCallsInRequest {
		t.Errorf("maxCallsInRequest is incorrect: got %v want %v", val, jmap.MaxCallsInRequest)
	}
}

func TestHealthCheck(t *testing.T) {
	app := newTestApplication(t)
	app.Port = 8443

	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.healthCheck)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "status: available\nport: 8443\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %q want %q",
			rr.Body.String(), expected)
	}
}

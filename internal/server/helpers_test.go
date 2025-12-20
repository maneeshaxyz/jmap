package server

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerError(t *testing.T) {
	var logBuf bytes.Buffer
	app := &Application{
		ErrorLog: log.New(&logBuf, "", 0),
	}

	rr := httptest.NewRecorder()
	app.serverError(rr, errors.New("test error"))

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("wrong status code: got %d, want %d", rr.Code, http.StatusInternalServerError)
	}

	if logBuf.Len() == 0 {
		t.Error("error log should not be empty")
	}

	expectedBody := http.StatusText(http.StatusInternalServerError) + "\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("unexpected body: got %q, want %q", rr.Body.String(), expectedBody)
	}
}

func TestClientError(t *testing.T) {
	app := newTestApplication(t)

	rr := httptest.NewRecorder()
	app.clientError(rr, http.StatusBadRequest)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("wrong status code: got %d, want %d", rr.Code, http.StatusBadRequest)
	}

	expectedBody := http.StatusText(http.StatusBadRequest) + "\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("unexpected body: got %q, want %q", rr.Body.String(), expectedBody)
	}
}

func TestNotFound(t *testing.T) {
	app := newTestApplication(t)

	rr := httptest.NewRecorder()
	app.notFound(rr)

	if rr.Code != http.StatusNotFound {
		t.Errorf("wrong status code: got %d, want %d", rr.Code, http.StatusNotFound)
	}

	expectedBody := http.StatusText(http.StatusNotFound) + "\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("unexpected body: got %q, want %q", rr.Body.String(), expectedBody)
	}
}

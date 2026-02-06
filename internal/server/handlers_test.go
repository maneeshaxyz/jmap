package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlers(t *testing.T) {
	app := &application{}
	mux := app.routes()

	tests := []struct {
		name           string
		method         string
		path           string
		body           string
		contentType    string
		expectedStatus int
	}{
		{
			name:           "GET healthcheck",
			method:         "GET",
			path:           "/healthcheck",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "GET JMAP session",
			method:         "GET",
			contentType:    "application/json",
			path:           "/.well-known/jmap",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "POST correct JMAP request",
			method:         "POST",
			path:           "/jmap/request",
			body:           `{"using":["urn:ietf:params:jmap:core"],"methodCalls":[]}`,
			contentType:    "application/json",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "POST JMAP request with no body",
			method:         "POST",
			path:           "/jmap/request",
			contentType:    "application/json",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "POST JMAP request with no content and no body",
			method:         "POST",
			path:           "/jmap/request",
			expectedStatus: http.StatusUnsupportedMediaType,
		},
		{
			name:           "Nonexistent path",
			method:         "GET",
			path:           "/nonexistent",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.body != "" {
				req = httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			} else {
				req = httptest.NewRequest(tt.method, tt.path, nil)
			}

			if tt.contentType != "" {
				req.Header.Set("Content-Type", tt.contentType)
			}

			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}
		})
	}
}

package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	app := &Application{}

	mux := app.routes()

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "GET home",
			method:         "GET",
			path:           "/",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "GET healthcheck",
			method:         "GET",
			path:           "/healthcheck",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "GET JMAP session",
			method:         "GET",
			path:           "/.well-known/jmap",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "POST JMAP request",
			method:         "POST",
			path:           "/jmap/request",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Method not allowed - POST to home",
			method:         "POST",
			path:           "/",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Not found",
			method:         "GET",
			path:           "/nonexistent",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}
		})
	}
}

package server

import (
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		method   string
		url      string
		wantCode int
	}{
		//{"Post Resource (Wrong Method)", http.MethodGet, "/post/resource", http.StatusMethodNotAllowed},
		{"Home", http.MethodGet, "/", http.StatusOK},

		{"Healthcheck", http.MethodGet, "/healthcheck", http.StatusOK},

		{"Valid Session Request", http.MethodGet, "/.well-known/jmap", http.StatusOK},
		{"Invalid Session Request", http.MethodPost, "/.well-known/jmap", http.StatusMethodNotAllowed},

		{"Valid Jmap Request", http.MethodPost, "/jmap/request", http.StatusOK},
		{"Invalid Jmap Request", http.MethodGet, "/jmap/request", http.StatusMethodNotAllowed},

		{"Invalid Path", http.MethodGet, "/invalid/path", http.StatusNotFound},
	}
}

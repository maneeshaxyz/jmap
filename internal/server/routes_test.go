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
		{"Home", "GET", "/", http.StatusOK},
		//{"Post Resource (Wrong Method)", "GET", "/post/resource", http.StatusMethodNotAllowed},
		{"Healthcheck", "GET", "/healthcheck", http.StatusOK},
		{"Session", "GET", "/.well-known/jmap", http.StatusOK},
		{"Session", "POST", "/jmap/request", http.StatusOK},
		{"Invalid Path", "GET", "/invalid/path", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var code int

			switch tt.method {
			case "GET":
				code, _, _ = ts.get(t, tt.url)
			case "POST":
				res, err := ts.Client().Post(ts.URL+tt.url, "application/json", nil)
				if err != nil {
					t.Fatal(err)
				}
				code = res.StatusCode
				res.Body.Close()
			default:
				t.Fatalf("method %s not supported in test", tt.method)
			}

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}
		})
	}
}

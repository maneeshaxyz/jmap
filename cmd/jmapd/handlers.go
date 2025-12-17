package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	if _, err := w.Write([]byte("Hello from your JMAP server")); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) getHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("This is my GET handler")); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Add("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	if _, err := w.Write([]byte("This is my POST handler")); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, "status: available"); err != nil {
		app.serverError(w, err)
		return
	}
	if _, err := fmt.Fprintf(w, "port: %d\n", app.port); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) wellKnownHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode("/jmap/session"); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) sessionHandler(w http.ResponseWriter, r *http.Request) {
	session := Session{
		Capabilities: Capabilities{
			"urn:ietf:params:jmap:core": CoreCapability{
				MaxSizeUpload:         50000000,
				MaxConcurrentUpload:   4,
				MaxSizeRequest:        10000000,
				MaxConcurrentRequests: 4,
				MaxCallsInRequest:     16,
				MaxObjectsInGet:       500,
				MaxObjectsInSet:       500,
				CollationAlgorithms:   []string{"i;ascii-numeric"},
			},
		},
		PrimaryAccounts: map[string]string{
			"urn:ietf:params:jmap:mail": "account1",
		},
		Username:       "user@example.com",
		APIURL:         "https://api.example.com/jmap/",
		DownloadURL:    "https://api.example.com/jmap/download/{accountId}/{blobId}/{type}/{name}",
		UploadURL:      "https://api.example.com/jmap/upload/{accountId}",
		EventSourceURL: "https://api.example.com/jmap/events?types={types}&closeafter={closeafter}&ping={ping}",
		State:          "some_state_string",
		Accounts:       make(Accounts), // Fill with actual accounts
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(session); err != nil {
		app.serverError(w, err)
		return
	}
}

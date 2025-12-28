package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maneeshaxyz/jmap/internal/jmap"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	if _, err := w.Write([]byte("Hello from your JMAP server")); err != nil {
		app.serverError(w, err)
		return
	}
}

// func (app *Application) getHandler(w http.ResponseWriter, r *http.Request) {
// 	if _, err := w.Write([]byte("This is my GET handler")); err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// }

// func (app *Application) postHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		w.Header().Add("Allow", "POST")
// 		app.clientError(w, http.StatusMethodNotAllowed)
// 		return
// 	}
// 	if _, err := w.Write([]byte("This is my POST handler")); err != nil {
// 		app.serverError(w, err)
// 		return
// 	}
// }

func (app *Application) healthCheck(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintln(w, "status: available"); err != nil {
		app.serverError(w, err)
		return
	}
	if _, err := fmt.Fprintf(w, "port: %d\n", app.Port); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *Application) sessionHandler(w http.ResponseWriter, r *http.Request) {
	session := jmap.BuildSession()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(session); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *Application) requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Add("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	var jmapReq jmap.JmapRequest
	err := json.NewDecoder(r.Body).Decode(&jmapReq)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(&jmapReq)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
}

// All handlers should:
// Accept a HTTP request, parse and validate it.
// Call some ServiceThing to do ImportantBusinessLogic with the data we get from step 1.
// Send an appropriate HTTP response depending on what ServiceThing returns.

package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	jmap "github.com/maneeshaxyz/jmap/internal/core"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	if _, err := w.Write([]byte("Hello from your JMAP server \n")); err != nil {
		app.serverError(w, err)
		return
	}
}

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
	session := jmap.NewSession()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(session); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *Application) requestHandler(w http.ResponseWriter, r *http.Request) {

	if isReqNotJson := app.isReqNotJson(r, w); isReqNotJson {
		return
	}

	var jmapReq jmap.JmapRequest
	err := json.NewDecoder(r.Body).Decode(&jmapReq)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// LOGIC based on methods

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&jmapReq)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
}

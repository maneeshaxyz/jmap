// All handlers should:
// 1 - Accept a HTTP request, parse and validate it.
// 2 - Call some ServiceThing to do ImportantBusinessLogic with the data we get from step 1.
// 3 - Send an appropriate HTTP response depending on what ServiceThing returns.

package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	jmap "github.com/maneeshaxyz/jmap/internal/core"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "available",
		"port":    strconv.Itoa(app.config.port),
		"version": app.config.version,
	}
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (app *application) sessionHandler(w http.ResponseWriter, r *http.Request) {
	session := jmap.NewSession()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(session); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) requestHandler(w http.ResponseWriter, r *http.Request) {
	if isReqNotJson := app.requireJSON(r, w); isReqNotJson {
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

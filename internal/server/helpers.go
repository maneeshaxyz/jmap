package server

import (
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"
)

func (app *Application) serverError(w http.ResponseWriter, err error) {
	app.logger.Error(
		"server error",
		slog.String("error", err.Error()),
		slog.String("stack", string(debug.Stack())),
	)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *Application) requireJSON(r *http.Request, w http.ResponseWriter) bool {
	if r.Header.Get("Content-Type") != "application/json" && !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json;") {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return true
	}
	return false
}

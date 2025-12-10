package main

import (
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

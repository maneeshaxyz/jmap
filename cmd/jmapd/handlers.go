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
	}
}

func (app *application) getHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("This is my GET handler")); err != nil {
		app.serverError(w, err)
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
	}
}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "port: %d\n", app.config.port)
}

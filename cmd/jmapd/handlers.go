package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from your JMAP server"))
}

func (app *application) getHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my GET handler"))
}

func (app *application) postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Add("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("This is my POST handler"))
}

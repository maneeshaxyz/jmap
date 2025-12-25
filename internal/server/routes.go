package server

import "net/http"

func (a *Application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/healthcheck", a.healthCheck)
	mux.HandleFunc("/.well-known/jmap", a.sessionHandler)

	return mux
}

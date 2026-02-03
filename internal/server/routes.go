package server

import "net/http"

func (a *application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthcheck", a.healthCheck)
	mux.HandleFunc("GET /.well-known/jmap", a.sessionHandler)
	mux.HandleFunc("POST /jmap/request", a.requestHandler)

	return mux
}

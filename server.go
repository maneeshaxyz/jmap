package main

import (
	"net/http"
)

type JMAPServer struct {
}

func (j *JMAPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

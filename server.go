package main

import (
	"net/http"
)

type JMAPServer struct {
}

func (j *JMAPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"methodResponses":[]}`))
}

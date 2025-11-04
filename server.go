package main

import (
	"encoding/json"
	"net/http"
)

type JMAPServer struct {
}

type JMAPResponse struct {
	MethodResponses [][]interface{} `json:"methodResponses"`
	SessionState    string          `json:"sessionState"`
}

func (j *JMAPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := JMAPResponse{
		MethodResponses: [][]interface{}{},
		SessionState:    "dummy",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

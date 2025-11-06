package main

import (
	"encoding/json"
	"net/http"
)

type JMAPServer struct {
}

type JMAPResponse struct {
	MethodResponses [][]any `json:"methodResponses"`
	SessionState    string  `json:"sessionState"`
}

type JMAPRequest struct {
	Using       []string `json:"using"`
	MethodCalls [][]any  `json:"methodCalls"`
}

func (j *JMAPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var req JMAPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := JMAPResponse{
		MethodResponses: [][]any{},
		SessionState:    "dummy",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

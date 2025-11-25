package main

import (
	"encoding/json"
	"net/http"
)

type JMAPRequest struct {
	Using       []string `json:"using"`
	MethodCalls [][]any  `json:"methodCalls"`
}

type JMAPResponse struct {
	MethodResponses [][]any `json:"methodResponses"`
	SessionState    string  `json:"sessionState"`
}

type JMAPServer struct{}

func (j *JMAPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var request JMAPRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	methodResponses := make([][]any, 0)

	for _, call := range request.MethodCalls {
		// Not a valid call, skip it
		if len(call) < 3 {
			continue
		}
		// Not a string then skip it
		methodName, ok := call[0].(string)
		if !ok {
			continue
		}

		if methodName == "Echo/echo" {
			methodResponses = append(methodResponses, call)
		}
	}

	response := JMAPResponse{
		MethodResponses: methodResponses,
		SessionState:    "dummy",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

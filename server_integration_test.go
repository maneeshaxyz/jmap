package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakingUserAndRetrievingThem(t *testing.T) {
	store := NewInMemoryUserStore()
	server := UserServer{store}
	user := "PunchiManike"

	server.ServeHTTP(httptest.NewRecorder(), newPostRequest(user))
	server.ServeHTTP(httptest.NewRecorder(), newPostRequest(user))
	server.ServeHTTP(httptest.NewRecorder(), newPostRequest(user))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetRequest(user))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "changed value")
}

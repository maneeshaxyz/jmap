package main

import (
	"fmt"
	"net/http"
	"strings"
)

type UserString interface {
	GetUserString(name string) string
	ChangeUserValues(value string)
}

type UserServer struct {
	store UserString
}

func GetUserString(name string) string {
	if name == "Piyaseli" {
		return "Piyaseli's Mailbox"
	}

	if name == "Siripala" {
		return "Siripala's Mailbox"
	}

	return ""
}

func (u *UserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user := strings.TrimPrefix(r.URL.Path, "/users/")

	switch r.Method {
	case http.MethodPost:
		u.postString(w, user)
	case http.MethodGet:
		u.getString(w, user)
	}
}

func (u *UserServer) getString(w http.ResponseWriter, user string) {
	userString := u.store.GetUserString(user)
	if userString == "" {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, u.store.GetUserString(user))
}

func (u *UserServer) postString(w http.ResponseWriter, user string) {
	u.store.ChangeUserValues(user)
	w.WriteHeader(http.StatusAccepted)
}

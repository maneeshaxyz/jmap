package main

import (
	"fmt"
	"net/http"
	"strings"
)

const port = ":6969"

func GetReq(w http.ResponseWriter, r *http.Request) {

	user := strings.TrimPrefix(r.URL.Path, "/users/")

	fmt.Fprint(w, GetUserString(user))
}

func GetUserString(name string) string {
	if name == "Piyaseli" {
		return "Piyaseli's mailbox"
	}

	if name == "Siripala" {
		return "Siripala's mailbox"
	}

	return ""
}

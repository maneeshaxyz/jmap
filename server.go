package main

import (
	"fmt"
	"net/http"
)

const port = ":6969"

func GetReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Siripala's mailbox")
}

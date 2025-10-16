package main

import (
	"log"
	"net/http"
)

func main() {
	handler := &UserServer{NewInMemoryUserStore()}
	log.Fatal(http.ListenAndServe(":8080", handler))
}

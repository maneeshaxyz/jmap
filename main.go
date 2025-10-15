package main

import (
	"log"
	"net/http"
)

type InMemoryUserStore struct{}

func (i *InMemoryUserStore) GetUserString(name string) string {
	return "123"
}

func (i *InMemoryUserStore) ChangeUserValues(name string) {}

func main() {
	handler := &UserServer{&InMemoryUserStore{}}
	log.Fatal(http.ListenAndServe(":8080", handler))
}

package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	server := &JMAPServer{}

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      server,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("JMAP server starting on port 8080")
	log.Fatal(srv.ListenAndServe())
}

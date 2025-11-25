package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("JMAP server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", &JMAPServer{}))
}

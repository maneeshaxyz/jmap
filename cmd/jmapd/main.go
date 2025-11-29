package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	port := flag.String("port", ":8080", "HTTP port")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/get/resource", getHandler)
	mux.HandleFunc("/post/resource", postHandler)

	log.Printf("Starting server on %s", *port)
	err := http.ListenAndServe(*port, mux)
	log.Fatal(err)
}

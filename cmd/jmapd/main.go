package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	port := flag.String("port", ":8080", "HTTP port")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/get/resource", getHandler)
	mux.HandleFunc("/post/resource", postHandler)

	infoLog.Printf("Starting server on %s", *port)
	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

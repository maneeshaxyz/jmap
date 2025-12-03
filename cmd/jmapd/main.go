package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port int
}

type application struct {
	config
	errorLog *log.Logger
	infoLog  *log.Logger
}

// main.go handles:
// - parsing the runtime configuration settings for the JMAP server,
// - init of loggers to be injected to handlers.go
// - Running the HTTP server

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "HTTP port")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:   cfg,
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     fmt.Sprintf(":%d", cfg.port),
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", srv.Addr)

	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

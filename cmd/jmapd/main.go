package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port     int
	certfile string
	keyfile  string
}

type application struct {
	config
	errorLog *log.Logger
	infoLog  *log.Logger
}

// main.go handles:
// - parsing the runtime configuration settings for the JMAP server,
// - init of loggers to be injected to handlers.go
// - Running the HTTPS server

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8443, "HTTP port")
	flag.StringVar(&cfg.certfile, "certfile", "server.crt", "Cert File")
	flag.StringVar(&cfg.keyfile, "keyfile", "server.key", "Key File")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if cfg.certfile == "" || cfg.keyfile == "" {
		errorLog.Fatal("TLS enabled: certfile and keyfile must be provided")
	}

	app := &application{
		config:   cfg,
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	infoLog.Printf("Starting server on %s", srv.Addr)

	err := srv.ListenAndServeTLS(cfg.certfile, cfg.keyfile)
	errorLog.Fatal(err)
}

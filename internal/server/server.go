package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Port     int
	Certfile string
	Keyfile  string
}

type Application struct {
	Config
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func Run() {
	var cfg Config

	flag.IntVar(&cfg.Port, "port", 8443, "HTTP port")
	flag.StringVar(&cfg.Certfile, "certfile", "server.crt", "Cert File")
	flag.StringVar(&cfg.Keyfile, "keyfile", "server.key", "Key File")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if cfg.Certfile == "" || cfg.Keyfile == "" {
		errorLog.Fatal("TLS enabled: certfile and keyfile must be provided")
	}

	app := &Application{
		Config:   cfg,
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	infoLog.Printf("Starting server on %s", srv.Addr)

	err := srv.ListenAndServeTLS(cfg.Certfile, cfg.Keyfile)
	errorLog.Fatal(err)
}

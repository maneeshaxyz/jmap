package server

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type config struct {
	port     int
	certfile string
	keyfile  string
	version  string
}

type application struct {
	config config
	logger *slog.Logger
}

func Run() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8443, "HTTP port")
	flag.StringVar(&cfg.certfile, "certfile", "server.crt", "Cert File")
	flag.StringVar(&cfg.keyfile, "keyfile", "server.key", "Key File")
	flag.StringVar(&cfg.version, "version", "v0.1.0", "Version")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if cfg.certfile == "" || cfg.keyfile == "" {
		logger.Error("TLS enabled: certfile and keyfile must be provided")
	}

	app := &application{
		config: cfg,
		logger: logger,
	}

	// addr expects a string in ":8080" format
	// TODO: need to refine timeouts
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("Starting server on", "addr", srv.Addr, "version", cfg.version)

	err := srv.ListenAndServeTLS(cfg.certfile, cfg.keyfile)
	logger.Error(err.Error())
	os.Exit(1)
}

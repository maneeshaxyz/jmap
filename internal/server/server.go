package server

import (
	"errors"
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
	flag.StringVar(&cfg.version, "version", "v0.1.0", "Version")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

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

	if err := srv.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logger.Info("server stopped", "reason", "shutdown")
			return
		}
		logger.Error("server stopped with error", "error", err)
		os.Exit(1)
	}
}

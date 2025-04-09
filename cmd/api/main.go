package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "1.0.0"

type config struct {
	port string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	cfg := config{
		port: os.Getenv("API_PORT"),
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.port),
		Handler: app.routes(),
	}

	logger.Printf("Starting server on port %s", srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}

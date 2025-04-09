package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "1.0.0"

type Config struct {
	port string
}

type Application struct {
	config Config
	logger *log.Logger
}

func main() {
	cfg := Config{
		port: os.Getenv("API_PORT"),
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
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

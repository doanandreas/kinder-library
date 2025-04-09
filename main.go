package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bruhhhh h h h"))
	})

	port := os.Getenv("API_PORT")
	log.Printf("API_PORT: %s\n", port)
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	log.Printf("Starting server on port :%s\n", port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

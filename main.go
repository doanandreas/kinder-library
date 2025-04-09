package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello KinderCastle!"))
	})

	log.Printf("Starting server on port %s\n", ":8080")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}

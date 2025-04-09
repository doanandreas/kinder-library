package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *Application) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/v1/healthcheck", app.healthcheckHandler)

	router.Get("/v1/books", app.listBooksHandler)
	router.Post("/v1/books", app.insertBooksHandler)
	router.Put("/v1/books/{id}", app.updateBooksHandler)
	router.Delete("/v1/books/{id}", app.deleteBooksHandler)

	return router
}

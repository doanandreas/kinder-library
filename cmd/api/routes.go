package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/v1/healthcheck", app.healthcheckHandler)

	router.Route("/v1/books", func(r chi.Router) {
		r.Use(app.authRequest)

		r.Get("/", app.listBooksHandler)
		r.Post("/", app.insertBooksHandler)
		r.Put("/{id}", app.updateBooksHandler)
		r.Delete("/{id}", app.deleteBooksHandler)
	})

	return router
}

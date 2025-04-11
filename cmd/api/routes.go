package main

import (
	_ "embed"
	"github.com/flowchartsman/swaggerui"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:embed docs/swagger.json
var spec []byte

func (app *Application) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/v1/healthcheck", app.healthcheckHandler)
	router.Mount("/v1/swagger", http.StripPrefix("/v1/swagger", swaggerui.Handler(spec)))

	router.Route("/v1/books", func(r chi.Router) {
		r.Use(app.authRequest)

		r.Get("/", app.listBooksHandler)
		r.Post("/", app.insertBooksHandler)
		r.Put("/{id}", app.updateBooksHandler)
		r.Delete("/{id}", app.deleteBooksHandler)
	})

	return router
}

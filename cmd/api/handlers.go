package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (app *Application) listBooksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List all books with pagination")
}

func (app *Application) insertBooksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Insert a book")
}

func (app *Application) updateBooksHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Update one book by ID: %d\n", id)
}

func (app *Application) deleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Delete one book by ID: %d\n", id)
}

func (app *Application) healthcheckHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "status: available\n")
	fmt.Fprintf(w, "version: %s\n", version)
}

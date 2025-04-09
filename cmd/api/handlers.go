package main

import (
	"fmt"
	"net/http"
)

func (app *Application) listBooksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List all books with pagination")
}

func (app *Application) insertBooksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Insert a book")
}

func (app *Application) updateBooksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update one book by ID")
}

func (app *Application) deleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete one book by ID")
}

func (app *Application) healthcheckHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "status: available\n")
	fmt.Fprintf(w, "version: %s\n", version)
}

package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (app *Application) listBooksHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	data := map[string]int{
		"page":      page,
		"page_size": pageSize,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Printf("Failed to write JSON response: %v\n", err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
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
	data := map[string]string{
		"status":  "available",
		"version": version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Printf("Failed to write JSON response: %v\n", err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

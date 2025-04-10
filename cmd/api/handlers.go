package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/doanandreas/kinder-library/internal/data"
	"github.com/go-chi/chi/v5"
)

func (app *Application) listBooksHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	res := data.BookListResponse{
		Metadata: data.Metadata{
			CurrentPage:  page,
			PageSize:     pageSize,
			FirstPage:    1,
			LastPage:     5,
			TotalRecords: 37,
		},
		Books: []data.Book{
			{
				ID:          1,
				Title:       "Let's Go!",
				Author:      "Alex Edwards",
				Pages:       426,
				Description: "Introduction REST API Golang",
				Rating:      4.62,
				Genres:      []string{"Programming", "Go", "Best-seller"},
			},
			{
				ID:          2,
				Title:       "Let's Go Further!",
				Author:      "Alex Edwards",
				Pages:       590,
				Description: "Advanced REST API Golang",
				Rating:      4.77,
				Genres:      []string{"Programming", "Go", "Best-seller"},
			},
		},
	}

	err := app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.logger.Printf("Failed to write JSON response: %v\n", err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (app *Application) insertBooksHandler(w http.ResponseWriter, r *http.Request) {
	book := data.BookResponse{
		Book: data.Book{
			ID:          1,
			Title:       "Let's Go Further!",
			Author:      "Alex Edwards",
			Pages:       590,
			Description: "Advanced REST API Golang",
			Rating:      4.77,
			Genres:      []string{"Programming", "Go", "Best-seller"},
		},
	}

	err := app.writeJSON(w, http.StatusOK, book, nil)
	if err != nil {
		app.logger.Printf("Failed to write JSON response: %v\n", err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (app *Application) updateBooksHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	book := data.BookResponse{
		Book: data.Book{
			ID:          int64(id),
			Title:       "Let's Go Further!",
			Author:      "Alex Edwards",
			Pages:       590,
			Description: "Advanced REST API Golang",
			Rating:      4.77,
			Genres:      []string{"Programming", "Go", "Best-seller"},
		},
	}

	err = app.writeJSON(w, http.StatusOK, book, nil)
	if err != nil {
		app.logger.Printf("Failed to write JSON response: %v\n", err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
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
	res := map[string]string{
		"status":  "available",
		"version": version,
	}

	err := app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.logger.Printf("Failed to write JSON response: %v\n", err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

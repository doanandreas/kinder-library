package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/doanandreas/kinder-library/internal/data"
	"github.com/doanandreas/kinder-library/internal/validator"

	"github.com/go-chi/chi/v5"
)

func (app *Application) listBooksHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	filters := data.Filters{
		Page:     page,
		PageSize: pageSize,
	}

	books, pagination, err := app.models.Books.List(filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	res := data.BookListResponse{
		Pagination: pagination,
		Books:      books,
	}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) insertBooksHandler(w http.ResponseWriter, r *http.Request) {
	var input data.BookRequest
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	input.Validate(v)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	book := &data.Book{
		Title:       input.Title,
		Author:      input.Author,
		Pages:       input.Pages,
		Description: input.Description,
		Rating:      input.Rating,
		Genres:      input.Genres,
	}

	err = app.models.Books.Insert(book)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	res := data.BookResponse{
		Book: *book,
	}

	err = app.writeJSON(w, http.StatusCreated, res, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) updateBooksHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var input data.BookRequest
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	input.Validate(v)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	book := &data.Book{
		ID:          int64(id),
		Title:       input.Title,
		Author:      input.Author,
		Pages:       input.Pages,
		Description: input.Description,
		Rating:      input.Rating,
		Genres:      input.Genres,
	}

	err = app.models.Books.Update(book)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	res := data.BookResponse{
		Book: *book,
	}

	err = app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) deleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Books.Delete(int64(id))
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusNoContent, nil, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"status":  "available",
		"version": version,
	}

	err := app.writeJSON(w, http.StatusOK, res, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

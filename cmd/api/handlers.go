package main

import (
	"errors"
	"github.com/doanandreas/kinder-library/internal/repository"
	"net/http"
	"strconv"

	"github.com/doanandreas/kinder-library/internal/data"
	"github.com/doanandreas/kinder-library/internal/validator"

	"github.com/go-chi/chi/v5"
)

func (app *Application) listBooksHandler(w http.ResponseWriter, r *http.Request) {
	input := &data.FiltersRequest{
		Page:     r.URL.Query().Get("page"),
		PageSize: r.URL.Query().Get("page_size"),
	}

	v := validator.New()
	input.Validate(v)
	if v.Valid() == false {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	filters := data.ParseFilters(input)
	filters.Validate(v)
	if v.Valid() == false {
		app.failedValidationResponse(w, r, v.Errors)
		return
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
	if v.Valid() == false {
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
		switch {
		case errors.Is(err, repository.ErrDuplicateTitle):
			v.AddError("title", "must be unique")
			app.failedValidationResponse(w, r, v.Errors)
			return
		default:
			app.serverErrorResponse(w, r, err)
		}

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
	v := validator.New()

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		v.AddError("id", "must be an integer")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	if id <= 0 {
		v.AddError("id", "must be a positive, non-zero integer")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	var input data.BookRequest
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	input.Validate(v)
	if v.Valid() == false {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	book := &data.Book{
		ID:          id,
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
		case errors.Is(err, repository.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		case errors.Is(err, repository.ErrDuplicateTitle):
			v.AddError("title", "must be unique")
			app.failedValidationResponse(w, r, v.Errors)
			return
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
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		v := validator.New()
		v.AddError("id", "must be an integer")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	if id <= 0 {
		v := validator.New()
		v.AddError("id", "must be a positive, non-zero integer")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Books.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
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

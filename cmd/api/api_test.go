package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/doanandreas/kinder-library/internal/data"
	"github.com/doanandreas/kinder-library/internal/mocks"
	"github.com/doanandreas/kinder-library/internal/repository"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Test_HealthCheck(t *testing.T) {
	tests := []struct {
		name   string
		status string
	}{
		{"API Available", "available"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var app Application

			// SUT == system under test
			sut := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/v1/healthcheck", nil)

			handler := http.HandlerFunc(app.healthcheckHandler)
			handler.ServeHTTP(sut, r)

			body, err := io.ReadAll(sut.Body)
			if err != nil {
				t.Fatal(err)
			}

			var js healthCheck
			err = json.Unmarshal(body, &js)
			if err != nil {
				t.Fatal(err)
			}

			if js.Status != "available" {
				t.Errorf("got %s; expected 'available'", js.Status)
			}
		})
	}
}

func Test_InsertBook(t *testing.T) {
	var wrongTitle string
	wrongTitle = "Existing Title"

	tests := []struct {
		name        string
		title       string
		pages       int
		rating      float64
		genres      []string
		eStatusCode int
	}{
		{"Book inserted", "Unique Title", 123, 4.56, []string{"testing", "mocking"}, http.StatusCreated},
		{"Book title already exists", wrongTitle, 123, 4.56, []string{"testing", "mocking"}, http.StatusUnprocessableEntity},
		{"Book title is empty", "", 123, 4.56, []string{"testing", "mocking"}, http.StatusUnprocessableEntity},
		{"Book pages are non-positive", "Unique Title", -2, 4.56, []string{"testing", "mocking"}, http.StatusUnprocessableEntity},
		{"Book rating is outside valid range", "Unique Title", -2, 7.12, []string{"testing", "mocking"}, http.StatusUnprocessableEntity},
		{"Book rating has too many decimals", "Unique Title", -2, 7.123456, []string{"testing", "mocking"}, http.StatusUnprocessableEntity},
		{"Book genres contains duplicates", "Unique Title", 123, 4.56, []string{"duplicate", "duplicate"}, http.StatusUnprocessableEntity},
		{"Book genres contains empty string", "Unique Title", 123, 4.56, []string{"testing", ""}, http.StatusUnprocessableEntity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := Application{
				models: &repository.Models{
					Books: mocks.InsertBookMock(wrongTitle),
				},
				logger: log.New(io.Discard, "", 0),
			}

			req := data.BookRequest{
				Title:       tt.title,
				Author:      "Test Author",
				Pages:       tt.pages,
				Description: "Just testing!",
				Rating:      tt.rating,
				Genres:      tt.genres,
			}

			body, err := json.Marshal(req)
			if err != nil {
				t.Fatal(err)
			}

			sut := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/v1/books", bytes.NewReader(body))

			handler := http.HandlerFunc(app.insertBooksHandler)
			handler.ServeHTTP(sut, r)

			if sut.Result().StatusCode != tt.eStatusCode {
				t.Errorf("got '%d'; expected '%d'", sut.Result().StatusCode, tt.eStatusCode)
			}
		})
	}
}

func Test_UpdateBook(t *testing.T) {
	var correctId int64
	var wrongTitle string
	correctId = 7
	wrongTitle = "Existing Title"

	tests := []struct {
		name        string
		id          string
		title       string
		eStatusCode int
	}{
		{"Book updated", strconv.Itoa(int(correctId)), "Unique Title", http.StatusOK},
		{"Book title is not unique", strconv.Itoa(int(correctId)), wrongTitle, http.StatusUnprocessableEntity},
		{"Mandatory field is empty", strconv.Itoa(int(correctId)), "", http.StatusUnprocessableEntity},
		{"Book ID not found", "2", "Unique Title", http.StatusNotFound},
		{"Book ID is zero", "0", "Dummy Title", http.StatusUnprocessableEntity},
		{"Book ID is negative", "-3", "Dummy Title", http.StatusUnprocessableEntity},
		{"Book ID is not int64", "hello", "Dummy Title", http.StatusUnprocessableEntity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := Application{
				models: &repository.Models{
					Books: mocks.UpdateBookMock(correctId, wrongTitle),
				},
				logger: log.New(io.Discard, "", 0),
			}

			req := data.BookRequest{
				Title:       tt.title,
				Author:      "Test Author",
				Pages:       123,
				Description: "Just testing!",
				Rating:      4.53,
				Genres:      []string{"testing", "mocking"},
			}

			body, err := json.Marshal(req)
			if err != nil {
				t.Fatal(err)
			}

			sut := httptest.NewRecorder()
			r := httptest.NewRequest("PUT", "/v1/movies/{id}", bytes.NewReader(body))

			rCtx := chi.NewRouteContext()
			rCtx.URLParams.Add("id", tt.id)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rCtx))

			handler := http.HandlerFunc(app.updateBooksHandler)
			handler.ServeHTTP(sut, r)

			if sut.Result().StatusCode != tt.eStatusCode {
				t.Errorf("got '%d'; expected '%d'", sut.Result().StatusCode, tt.eStatusCode)
			}
		})
	}
}

func Test_DeleteBook(t *testing.T) {
	var correctId int64
	correctId = 7

	tests := []struct {
		name        string
		id          string
		eStatusCode int
	}{
		{"Book deleted", strconv.Itoa(int(correctId)), http.StatusNoContent},
		{"Book ID not found", "2", http.StatusNotFound},
		{"Book ID is zero", "0", http.StatusUnprocessableEntity},
		{"Book ID is negative", "-3", http.StatusUnprocessableEntity},
		{"Book ID is not int64", "hello", http.StatusUnprocessableEntity},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := Application{
				models: &repository.Models{
					Books: mocks.DeleteBookMock(correctId),
				},
				logger: log.New(io.Discard, "", 0),
			}

			sut := httptest.NewRecorder()
			r := httptest.NewRequest("DELETE", "/v1/movies/{id}", nil)

			rCtx := chi.NewRouteContext()
			rCtx.URLParams.Add("id", tt.id)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rCtx))

			handler := http.HandlerFunc(app.deleteBooksHandler)
			handler.ServeHTTP(sut, r)

			if sut.Result().StatusCode != tt.eStatusCode {
				t.Errorf("got '%d'; expected '%d'", sut.Result().StatusCode, tt.eStatusCode)
			}
		})
	}
}

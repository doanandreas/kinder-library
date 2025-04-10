package main

import (
	"context"
	"encoding/json"
	"github.com/doanandreas/kinder-library/internal/mocks"
	"github.com/doanandreas/kinder-library/internal/repository"
	"github.com/go-chi/chi/v5"
	"io"
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

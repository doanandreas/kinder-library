package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
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

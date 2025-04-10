package main

import (
	"net/http"
)

type healthCheck struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := healthCheck{
		Status:  "available",
		Version: version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

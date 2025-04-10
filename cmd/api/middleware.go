package main

import "net/http"

func (app *Application) authRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "Basic ZG9hbjpkaWRpbmRpbmc=" {
			app.unauthorizedResponse(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

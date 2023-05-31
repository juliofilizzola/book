package middlewares

import (
	"api/cmd/auth"
	"api/cmd/response"
	"fmt"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s, %s, %s\n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidToken(r); err != nil {
			response.Err(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}

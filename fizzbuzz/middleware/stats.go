package middleware

import (
	"net/http"
	"strings"
)

func StatsMux(mux *http.ServeMux, hook string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Index(r.URL.Path, hook) != -1 {
			// Store api calls.
		}

		mux.ServeHTTP(w, r)
	}
}

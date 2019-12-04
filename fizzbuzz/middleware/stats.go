package middleware

import (
	"github.com/jkuma/gophercises/fizzbuzz/repository"
	"log"
	"net/http"
	"strings"
)

func StatsMux(mux *http.ServeMux, hook string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Index(r.URL.Path, hook) != -1 {
			// Store api calls.
			err := repository.Update(r)

			if err != nil {
				log.Fatal(err)
			}
		}

		mux.ServeHTTP(w, r)
	}
}

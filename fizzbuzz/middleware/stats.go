package middleware

import (
	"github.com/jkuma/gophercises/fizzbuzz/repository"
	"log"
	"net/http"
	"strings"
	"sync"
)

type apiCall struct {
	req *http.Request
	m   *sync.Mutex
}

// StatsMux is a middleware to be used as a server handler. The current
// request is processed whereas the hook string is contained in its
// fully qualified uri. If so, this call is stored and its counter is
// incremented.
func StatsMux(mux *http.ServeMux, hook string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Index(r.URL.Path, hook) != -1 {
			go storeApiCall(apiCall{req: r, m: &sync.Mutex{}})
		}

		mux.ServeHTTP(w, r)
	}
}

func storeApiCall(call apiCall) {
	call.m.Lock()
	defer call.m.Unlock()
	err := repository.Update(call.req)

	if err != nil {
		log.Fatal(err)
	}
}

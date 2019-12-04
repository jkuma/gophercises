package api

import (
	"encoding/binary"
	"encoding/json"
	"github.com/jkuma/gophercises/fizzbuzz/repository"
	"net/http"
	"net/url"
)

func GetStats(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		type response struct {
			Uri        string
			Hits       int
			Parameters url.Values
		}

		u, h, p, err := getTopApiCall()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		js, err := json.Marshal(response{Uri: u, Hits: h, Parameters: p})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func getTopApiCall() (uri string, hits int, parameters url.Values, err error) {
	var key []byte
	key, err = repository.FetchHighestScore()

	if err == nil {
		val, err := repository.Get(key)
		u, err := url.Parse(string(key))

		return uri, int(binary.BigEndian.Uint64(val)), u.Query(), err
	}

	return uri, hits, parameters, err
}

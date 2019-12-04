package api

import 	(
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

		u, h, p, err := stsResults()
		httpErr(w, err)

		js, err := json.Marshal(response{Uri: u, Hits: h, Parameters: p})
		httpErr(w, err)

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func stsResults() (uri string, hits int, parameters url.Values, err error) {
	var key []byte
	key, hits, err = repository.HighScore()

	if err == nil {
		uri = string(key)

		if u, err := url.Parse(uri); err == nil {
			parameters = u.Query()
		}
	}

	return uri, hits, parameters, err
}

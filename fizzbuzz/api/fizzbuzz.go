package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type response struct {
	Uri     string
	Results []string
}

func GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		js, err := json.Marshal(response{
			Uri:     getUri(r),
			Results: getResults(parseQueryParameters(r.URL.Query())),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func parseQueryParameters(query url.Values) (int, int, int, string, string) {
	var int1, int2, limit int = 3, 5, 100
	var str1, str2 string = "fizz", "buzz"

	for n, v := range query {
		switch n {
		case "int1":
			i, _ := strconv.Atoi(v[0])
			int1 = i
		case "int2":
			i, _ := strconv.Atoi(v[0])
			int2 = i
		case "limit":
			i, _ := strconv.Atoi(v[0])
			limit = i
		case "str1":
			str1 = v[0]
		case "str2":
			str2 = v[0]
		}
	}

	return int1, int2, limit, str1, str2
}

func getResults(int1, int2, limit int, str1, str2 string) []string {
	var results = make([]string, limit)

	for k, _ := range results {
		if (k+1)%int1 == 0 && (k+1)%int2 == 0 {
			results[k] = str1 + str2
			continue
		}

		if (k+1)%int1 == 0 {
			results[k] = str1
			continue
		}

		if (k+1)%int2 == 0 {
			results[k] = str2
			continue
		}

		results[k] = strconv.Itoa(k + 1)
	}

	return results
}

func getUri(r *http.Request) string {
	if !r.URL.IsAbs() {
		scheme := "http://"
		if r.TLS != nil {
			scheme = "https://"
		}

		return scheme + r.Host + r.URL.RequestURI()
	}

	return r.URL.RequestURI()
}

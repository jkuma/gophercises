package api

import (
	"encoding/json"
	"github.com/jkuma/gophercises/fizzbuzz/request"
	"net/http"
	"net/url"
	"strconv"
)

func GetFizzBuzz(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		type response struct {
			Uri     string
			Results []string
		}

		js, err := json.Marshal(response{
			Uri:     request.GetUri(r),
			Results: fbResults(fbQueryParams(r.URL.Query())),
		})

		httpErr(w, err)

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func fbQueryParams(query url.Values) (int1 int, int2 int, limit int, str1 string, str2 string) {
	int1, int2, limit = 3, 5, 100
	str1, str2 = "fizz", "buzz"

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

func fbResults(int1, int2, limit int, str1, str2 string) []string {
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

package request

import "net/http"

func GetUri(r *http.Request) string {
	if !r.URL.IsAbs() {
		scheme := "http://"
		if r.TLS != nil {
			scheme = "https://"
		}

		return scheme + r.Host + r.URL.RequestURI()
	}

	return r.URL.String()
}

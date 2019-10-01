package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type DefaultHandler struct {}

func (DefaultHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//Does it match a known arc?
	chapters := parseChapters()
	arc := strings.TrimLeft(req.URL.Path, "/")

	if v, ok := chapters[arc]; ok {
		t, err := template.New(arc).Parse(chapterTemplate())
		checkErr(err)

		err = t.Execute(rw, v)
		checkErr(err)
		return
	}

	//Lost in space? Lets try to redirect to intro arc.
	if _, ok := chapters["intro"]; ok{
		http.Redirect(rw, req, "/intro", http.StatusPermanentRedirect)
		return
	}

	//Something wrongs happened.
	rw.WriteHeader(http.StatusForbidden)
}

func parseChapters() map[string]Chapter {
	j, err := ioutil.ReadFile("scenario.json")
	checkErr(err)

	var chapters map[string]Chapter

	err = json.Unmarshal(j, &chapters)
	checkErr(err)

	return chapters
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", defaultMux())
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", DefaultHandler{})
	return mux
}

func chapterTemplate() string {
	return `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		<article>{{.Story}}</article>
		<h2>What will you do now?</h2>
		<ul>
		{{range .Options}}
			<li>
			<a href="/{{.Arc}}">{{.Text}}</a>
			</li>
		{{else}}<a href="/intro">Restart?</a></li>
		{{end}}
		</ul>
	</body>
</html>`
}

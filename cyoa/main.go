package main

import (
	"encoding/json"
	"flag"
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

type chaptersHandler struct {
	tpl      *template.Template
	chapters map[string]Chapter
}

var defaultTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		<article>
	    {{range .Story}}<p>{{.}}</p>{{end}}
		</article>
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

func main() {
	filename := flag.String("filename", "scenario.json", "Please enter the full path of json file.")
	port := flag.String("port", "8080", "The web server port.")
	flag.Parse()

	handler := chaptersHandler{
		tpl:      template.Must(template.New("").Parse(defaultTemplate)),
		chapters: parseChapters(*filename),
	}

	fmt.Println("Starting the server on :", *port)
	http.ListenAndServe(":"+*port, defaultMux(handler))
}

func (h chaptersHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//Does it match a known arc?
	arc := strings.TrimLeft(req.URL.Path, "/")

	if v, ok := h.chapters[arc]; ok {
		err := h.tpl.Execute(rw, v)
		checkErr(err)
		return
	}

	//Lost in space? Lets try to redirect to intro arc.
	if _, ok := h.chapters["intro"]; ok {
		http.Redirect(rw, req, "/intro", http.StatusPermanentRedirect)
		return
	}

	//Something wrongs happened.
	rw.WriteHeader(http.StatusForbidden)
}

func parseChapters(filename string) map[string]Chapter {
	j, err := ioutil.ReadFile(filename)
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

func defaultMux(handler chaptersHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	return mux
}

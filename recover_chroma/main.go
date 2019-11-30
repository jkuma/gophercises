package main

import (
	"fmt"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/jkuma/gophercises/recover_chroma/dbghtml"
	"github.com/jkuma/gophercises/recover_chroma/dbghtml/fcnt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/debug/", debugRoute)
	log.Fatal(http.ListenAndServe(":3000", recoverMux(mux)))
}

func debugRoute(rw http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[6:]
	f, err := fcnt.GetFileContent(filename)
	logError(err)

	number, err := strconv.Atoi(r.URL.Query().Get("highlight"))
	logError(err)

	highlight(rw, string(f.Content), number)
}

func recoverMux(mux *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("recover from", r)
				log.Println(string(debug.Stack()))
				w.Write(dbghtml.Html(debug.Stack()))
			}
		}()

		mux.ServeHTTP(w, r)
	}
}

func panicDemo(rw http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}

func panicAfterDemo(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "<h1>Hello!</h1>")
}

// Highlight some text using chroma handlers.
func highlight(w io.Writer, source string, number int) error {
	// Determine lexer.
	l := lexers.Get("go")
	if l == nil {
		l = lexers.Fallback
	}

	l = chroma.Coalesce(l)

	// Determine formatter.
	var i [][2]int
	i = append(i, [2]int{number, number})
	f := html.New(html.WithLineNumbers(), html.LineNumbersInTable(), html.HighlightLines(i))

	// Determine style.
	s := styles.Get("dracula")

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return f.Format(w, s, it)
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

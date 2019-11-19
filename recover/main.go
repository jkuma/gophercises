package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime/debug"
)

func main() {
	mux := http.NewServeMux()
	develop := os.Getenv("ENV") == "develop"
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":3000", recoverMux(mux, develop)))
}

func recoverMux(mux *http.ServeMux, develop bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("recover from", r)
				log.Println(string(debug.Stack()))
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Something went wrong"))

				if develop {
					w.Write(debug.Stack())
				}
			}
		}()

		rw := responseWriter{ResponseWriter: w}
		mux.ServeHTTP(&rw, r)
		rw.flush()
	}

}

type responseWriter struct {
	http.ResponseWriter
	writes [][]byte
	status int
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.writes = append(rw.writes, b)
	return len(b), nil
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
}

func (rw *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := rw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("the ResponseWriter does not support the Hijacker interface")
	}
	return hijacker.Hijack()
}

func (rw *responseWriter) Flush() {
	flusher, ok := rw.ResponseWriter.(http.Flusher)
	if !ok {
		return
	}
	flusher.Flush()
}

func (rw *responseWriter) flush() error {
	if rw.status != 0 {
		rw.ResponseWriter.WriteHeader(rw.status)
	}
	for _, write := range rw.writes {
		_, err := rw.ResponseWriter.Write(write)
		if err != nil {
			return err
		}
	}
	return nil
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	funcThatPanics()
}


func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

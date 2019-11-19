package main

import (
	"fmt"
	"log"
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

		mux.ServeHTTP(w, r)
	}

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

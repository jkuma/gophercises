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
	mux.HandleFunc("/panic/", panicDemo)
	mux.HandleFunc("/panic-after/", panicAfterDemo)
	mux.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func panicDemo(w http.ResponseWriter, r *http.Request) {
	defer funcThatRecorvers(w)
	funcThatPanics()
}

func funcThatRecorvers(w http.ResponseWriter) {
	if r := recover(); r != nil {
		log.Println("recover from", r)
		log.Println(string(debug.Stack()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))

		if os.Getenv("ENV") == "develop" {
			w.Write(debug.Stack())
		}
	}
}

func panicAfterDemo(w http.ResponseWriter, r *http.Request) {
	defer funcThatRecorvers(w)
	fmt.Fprint(w, "<h1>Hello!</h1>")
	funcThatPanics()
}

func funcThatPanics() {
	panic("Oh no!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

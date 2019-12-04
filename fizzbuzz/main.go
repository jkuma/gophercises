package main

import (
	"flag"
	"fmt"
	"github.com/jkuma/gophercises/fizzbuzz/api"
	"github.com/jkuma/gophercises/fizzbuzz/database"
	"github.com/jkuma/gophercises/fizzbuzz/middleware"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "3000", "Provides the port number for http server.")
	dbpath := flag.String("dbpath", "./fizzbuzz.db", "Provides the full qualified path to badger db.")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/fizzbuzz", api.GetFizzBuzz)
	mux.HandleFunc("/api/stats", api.GetStats)

	must(database.Init(*dbpath))
	must(http.ListenAndServe(":"+*port, middleware.StatsMux(mux, "/api")))
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

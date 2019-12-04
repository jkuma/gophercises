package main

import (
	"fmt"
	"github.com/jkuma/gophercises/fizzbuzz/api"
	"github.com/jkuma/gophercises/fizzbuzz/database"
	"github.com/jkuma/gophercises/fizzbuzz/middleware"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/fizzbuzz", api.GetFizzBuzz)
	mux.HandleFunc("/api/stats", api.GetStats)

	must(database.Init("./fizzbuzz.db"))
	must(http.ListenAndServe(":3000", middleware.StatsMux(mux, "/api")))
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

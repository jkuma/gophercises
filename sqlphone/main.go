package main

import "github.com/jkuma/gophercises/sqlphone/database"

func main() {
	db := database.DB()

	println(db.Ping())
}

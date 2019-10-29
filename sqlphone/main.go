package main

import (
	"github.com/jkuma/gophercises/sqlphone/database"
	"github.com/jkuma/gophercises/sqlphone/normalizer"
)

func main() {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query("SELECT * from phone_numbers")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var numberId int
		var number string

		if err := rows.Scan(&numberId, &number); err == nil {
			normalize := normalizer.Normalize(number)

			if database.NumberExists(normalize) {
				_ = database.DeleteNumber(normalize)
			}

			_ = database.UpdateNumber(numberId, normalize)
		}
	}

	if err := rows.Close(); err != nil {
		panic(err)
	}
}


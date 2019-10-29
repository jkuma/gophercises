package database

import "log"


func init() {
	checkerr(dropTable())
	checkerr(createTable())
	checkerr(setTable())
}

func dropTable() error {
	db := DB()

	tx, err := db.Begin()
	checkerr(err)
	defer tx.Rollback()

	_, err = tx.Exec("DROP TABLE IF EXISTS phone_numbers")

	checkerr(err)
	return tx.Commit()
}

func createTable() error {
	db := DB()

	tx, err := db.Begin()
	checkerr(err)
	defer tx.Rollback()

	_, err = tx.Exec("CREATE TABLE phone_numbers(number_id serial PRIMARY KEY, number VARCHAR (50) UNIQUE NOT NULL)")

	checkerr(err)
	return tx.Commit()
}

func setTable() error {
	db := DB()

	tx, err := db.Begin()
	checkerr(err)
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO phone_numbers(number) VALUES($1)")
	checkerr(err)
	defer stmt.Close()

	for _, n := range numbers() {
		if _, err := stmt.Exec(n); err != nil {
			log.Fatal(err)
		}
	}

	return tx.Commit()
}

func numbers() []string {
	return []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
}

func checkerr (err error) {
	if err != nil {
		log.Fatal(err)
	}
}
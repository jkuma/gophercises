package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 54320
	user     = "admin"
	password = "admin"
	dbname   = "phones"
)

func DB() *sql.DB {
	db, err := sql.Open("postgres", dsn())

	if err != nil {
		log.Fatalf("Cannot open database: %v", err)
	}

	return db
}

func dsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

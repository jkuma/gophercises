package database

import (
	bolt "go.etcd.io/bbolt"
	"log"
)

func DB(file string) (*bolt.DB, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(file, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
